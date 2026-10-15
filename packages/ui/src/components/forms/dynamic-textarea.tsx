// fields/dynamic-textarea.tsx
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/lib/utils";
import type { FormField } from "./form-builder";
import { DynamicFieldWrapper } from "./dynamic-field-wrapper";
import type { FieldValues } from "react-hook-form";

interface DynamicTextareaProps<TFieldValues extends FieldValues = any> {
  field: FormField<TFieldValues>;
  value?: string | number | readonly string[];
  onChange: (value: unknown) => void;
  error?: React.ReactNode;
}

export function DynamicTextarea({
  field,
  value,
  onChange,
  error,
}: DynamicTextareaProps) {
  return (
    <DynamicFieldWrapper
      name={field.name}
      label={field.label}
      desc={field.desc}
      error={error}
      labelClassName={field.classNames?.label}
      descClassName={field.classNames?.desc}
      errorClassName={field.classNames?.error}
    >
      <Textarea
        value={value ?? ""}
        onChange={(event) => onChange(event.target.value)}
        className={cn(field.classNames?.input, error && "border-destructive")}
      />
    </DynamicFieldWrapper>
  );
}
