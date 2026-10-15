// dynamic-form.tsx
import type { ReactNode } from "react";
import type {
  FieldValues,
  SubmitHandler,
  UseFormReturn,
} from "react-hook-form";
import { useDirection } from "@/components/ui/direction";

import { FormFieldRenderer } from "./form-field-renderer";
import type { FormField } from "./form-builder";

export interface DynamicFormProps<TFieldValues extends FieldValues> {
  form: UseFormReturn<TFieldValues> & {
    fields: FormField<TFieldValues>[];
    onSubmit: SubmitHandler<TFieldValues>;
  };
  //fields: FormField<TFieldValues>[];

  children?: ReactNode;

  className?: string;
}

export function DynamicForm<TFieldValues extends FieldValues>({
  form,
  children,
  className,
}: DynamicFormProps<TFieldValues>) {
  const dir = useDirection();
  return (
    <form
      dir={dir}
      onSubmit={form.handleSubmit(form.onSubmit)}
      className={className}
    >
      <div className="grid grid-cols-12 gap-4">
        {form.fields.map((field) => (
          <FormFieldRenderer key={field.name} field={field} form={form} />
        ))}
      </div>

      {children}
    </form>
  );
}
