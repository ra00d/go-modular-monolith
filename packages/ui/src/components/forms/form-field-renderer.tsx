// form-field-renderer.tsx
import { cn } from "@/lib/utils";
import type { FieldValues, UseFormReturn } from "react-hook-form";
import { Controller } from "react-hook-form";

import { DynamicCheckbox } from "./dynamic-checkbox";
import { DynamicInput } from "./dynamic-input";
import { DynamicSelect } from "./dynamic-select";
import { DynamicTextarea } from "./dynamic-textarea";
import type { FormField } from "./form-builder";

interface FormFieldRendererProps<TFieldValues extends FieldValues> {
  field: FormField<TFieldValues>;
  form: UseFormReturn<TFieldValues>;
}

export function FormFieldRenderer<TFieldValues extends FieldValues>({
  field,
  form,
}: FormFieldRendererProps<TFieldValues>) {
  const {
    name,
    type = "text",
    error,
    variant = "default",
    rowsSpan = 1,
    colsSpan = 12,
    onValueChange,
  } = field;

  if (type === "hidden") {
    return <input type="hidden" {...form.register(name)} />;
  }

  return (
    <div
      className={cn(
        `col-span-${colsSpan}`,
        rowsSpan > 1 && `row-span-${rowsSpan}`,
        variant === "contained" && "rounded-lg border bg-muted/30 p-4",
      )}
    >
      <Controller
        name={name}
        control={form.control}
        render={({ field: controllerField, fieldState }) => {
          const fieldError = error ?? fieldState.error?.message;

          const handleChange = (value: unknown) => {
            controllerField.onChange(value);
            onValueChange?.(value);
          };

          switch (type) {
            case "textarea":
              return (
                <DynamicTextarea
                  field={field}
                  value={controllerField.value}
                  onChange={handleChange}
                  error={fieldError}
                />
              );

            case "select":
              return (
                <DynamicSelect
                  field={field}
                  value={controllerField.value}
                  onChange={handleChange}
                  error={fieldError}
                />
              );

            case "checkbox":
              return (
                <DynamicCheckbox
                  field={field}
                  value={controllerField.value}
                  onChange={handleChange}
                  error={fieldError}
                />
              );

            default:
              return (
                <DynamicInput
                  field={field}
                  value={controllerField.value}
                  onChange={handleChange}
                  error={fieldError}
                  type={type}
                />
              );
          }
        }}
      />
    </div>
  );
}
