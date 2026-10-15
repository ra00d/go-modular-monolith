import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";

import { DynamicFieldWrapper } from "./dynamic-field-wrapper";
import type { FormField } from "./form-builder";
import type { FieldValues } from "react-hook-form";
import { useState } from "react";
import { InputGroup, InputGroupAddon, InputGroupInput } from "../ui";
import { Eye, EyeClosed, EyeOff } from "lucide-react";

interface DynamicInputProps<TFieldValues extends FieldValues = any> {
  field: FormField<TFieldValues>;
  value?: string | number | readonly string[];
  onChange: (value: unknown) => void;
  error?: React.ReactNode;
  type?: string;
}

export function DynamicInput({
  field,
  value,
  onChange,
  error,
  type = "text",
}: DynamicInputProps) {
  const isPassword = type === "password";
  const [passwordShown, setPasswordShown] = useState(false);
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
      <InputGroup>
        <InputGroupInput
          type={isPassword && passwordShown ? "text" : type}
          value={value ?? ""}
          onChange={(event) => {
            const raw = event.target.value;

            onChange(
              type === "number" ? (raw === "" ? undefined : Number(raw)) : raw,
            );
          }}
          className={cn(field.classNames?.input, error && "border-destructive")}
        />
        {isPassword ? (
          <InputGroupAddon
            align={"inline-end"}
            className="cursor-pointer"
            onClick={() => {
              setPasswordShown((prev) => !prev);
            }}
          >
            {passwordShown ? <EyeOff /> : <Eye />}
          </InputGroupAddon>
        ) : (
          <></>
        )}
      </InputGroup>
    </DynamicFieldWrapper>
  );
}
