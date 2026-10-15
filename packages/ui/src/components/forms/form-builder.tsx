// form.types.ts
import { zodResolver } from "@hookform/resolvers/zod";
import type { ReactNode } from "react";
import { useMemo } from "react";
import {
  useForm,
  type DefaultValues,
  type FieldValues,
  type Path,
  type SubmitHandler,
  type UseFormProps,
  type UseFormReturn,
} from "react-hook-form";
import { core, z, type _ZodType } from "zod";
type ZodTrait = {
  _zod: {
    def: any;
    [k: string]: any;
  };
};

export type FormFieldOption = {
  label: ReactNode;
  value: string | number | boolean;
  disabled?: boolean;
};

export interface FormField<T extends FieldValues = FieldValues> {
  name: Path<T>;
  type?:
    | "text"
    | "email"
    | "password"
    | "number"
    | "tel"
    | "url"
    | "date"
    | "datetime-local"
    | "textarea"
    | "select"
    | "checkbox"
    | "hidden";

  required?: boolean;
  label?: ReactNode;
  placeholder?: string;
  desc?: ReactNode;
  error?: ReactNode;

  variant?: "default" | "contained";

  options?: FormFieldOption[];

  defaultValue?: unknown;

  validation?: any;

  classNames?: {
    label?: string;
    error?: string;
    desc?: string;
    input?: string;
    select?: string;
  };

  rowsSpan?: number;
  colsSpan?: number;

  onValueChange?: (value: unknown) => void;
}

export const COL_SPANS: Record<number, string> = {
  1: "col-span-1",
  2: "col-span-2",
  3: "col-span-3",
  4: "col-span-4",
  5: "col-span-5",
  6: "col-span-6",
  7: "col-span-7",
  8: "col-span-8",
  9: "col-span-9",
  10: "col-span-10",
  11: "col-span-11",
  12: "col-span-12",
};

export const ROW_SPANS: Record<number, string> = {
  1: "row-span-1",
  2: "row-span-2",
  3: "row-span-3",
  4: "row-span-4",
};
// use-dynamic-form.ts
interface UseDynamicFormOptions<
  TFieldValues extends FieldValues = any,
> extends Omit<UseFormProps<TFieldValues>, "resolver"> {
  fields: FormField<TFieldValues>[];
  onSubmit: SubmitHandler<TFieldValues>;
}
export function useDynamicForm<TFieldValues extends FieldValues = FieldValues>({
  fields,
  defaultValues,
  onSubmit,
  ...options
}: UseDynamicFormOptions<TFieldValues>): UseFormReturn<TFieldValues> & {
  fields: FormField<TFieldValues>[];

  onSubmit: SubmitHandler<TFieldValues>;
} {
  const schema = useMemo(() => {
    const shape: Record<string, unknown> = {};

    for (const field of fields) {
      shape[field.name] = field.validation ?? z.any();
    }

    return z.object(shape) as any;
  }, [fields]);

  const resolvedDefaults = useMemo(() => {
    const values: Record<string, unknown> = {
      ...(defaultValues ?? {}),
    };

    for (const field of fields) {
      if (
        field.defaultValue !== undefined &&
        values[field.name] === undefined
      ) {
        values[field.name] = field.defaultValue;
      }
    }

    return values as DefaultValues<TFieldValues>;
  }, [fields, defaultValues]);

  const form = useForm<TFieldValues>({
    ...options,
    defaultValues: resolvedDefaults,
    resolver: zodResolver(schema),
  });
  return { ...form, fields, onSubmit };
}
