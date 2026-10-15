// fields/dynamic-checkbox.tsx
import { Checkbox } from "@/components/ui/checkbox";
import { cn } from "@/lib/utils";
import type { FormField } from "./form-builder";
import { DynamicFieldWrapper } from "./dynamic-field-wrapper";
import type { FieldValues } from "react-hook-form";
import { FieldDescription, FieldLabel } from "../ui";

interface DynamicCheckboxProps<TFieldValues extends FieldValues = any> {
  field: FormField<TFieldValues>;
  value?: unknown;
  onChange: (value: unknown) => void;
  error?: React.ReactNode;
}

export function DynamicCheckbox({
  field,
  value,
  onChange,
  error,
}: DynamicCheckboxProps) {
  return (
    <DynamicFieldWrapper
      name={field.name}
      error={error}
      errorClassName={field.classNames?.error}
    >
      <div className="flex items-center gap-2">
        <Checkbox
          checked={Boolean(value)}
          onCheckedChange={onChange}
          className={cn(field.classNames?.input)}
        />

        {field.label && (
          <FieldLabel
            htmlFor={field.name}
            className={cn("", field.classNames?.label)}
          >
            {field.label}
          </FieldLabel>
        )}
      </div>

      {field.desc && (
        <FieldDescription className={cn("", field.classNames?.desc)}>
          {field.desc}
        </FieldDescription>
      )}
    </DynamicFieldWrapper>
  );
}
