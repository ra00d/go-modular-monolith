import type { ReactNode } from "react";
import { cn } from "@/lib/utils";
import { Field, FieldDescription, FieldError, FieldLabel } from "../ui";

interface DynamicFieldWrapperProps {
  label?: ReactNode;
  desc?: ReactNode;
  error?: ReactNode;
  name: string;

  labelClassName?: string;
  descClassName?: string;
  errorClassName?: string;

  children: ReactNode;
}

export function DynamicFieldWrapper({
  label,
  desc,
  name,
  error,
  labelClassName,
  descClassName,
  errorClassName,
  children,
}: DynamicFieldWrapperProps) {
  return (
    <Field className="space-y-2">
      {label && (
        <FieldLabel htmlFor={name} className={cn("", labelClassName)}>
          {label}
        </FieldLabel>
      )}

      {children}

      {desc && (
        <FieldDescription className={cn(descClassName)}>
          {desc}
        </FieldDescription>
      )}

      {error && (
        <FieldError className={cn("", errorClassName)}>{error}</FieldError>
      )}
    </Field>
  );
}
