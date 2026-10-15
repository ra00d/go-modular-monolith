// fields/dynamic-select.tsx

import React from "react";
import type { FieldValues } from "react-hook-form";
import {
  Combobox,
  ComboboxContent,
  ComboboxEmpty,
  ComboboxInput,
  ComboboxItem,
  ComboboxList,
} from "../ui";
import { DynamicFieldWrapper } from "./dynamic-field-wrapper";
import type { FormField } from "./form-builder";

interface DynamicSelectProps<TFieldValues extends FieldValues = any> {
  field: FormField<TFieldValues>;
  value: unknown;
  onChange: (value: unknown) => void;
  error?: React.ReactNode;
}

export function DynamicSelect({
  field,
  value,
  onChange,
  error,
}: DynamicSelectProps) {
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
      <Combobox
        // autoHighlight
        //value={value == null ? undefined : String(value)}
        onValueChange={(value) => {
          const option = field.options?.find(
            (option) => String(option.value) === value.value,
          );

          onChange(option?.value ?? value);
        }}
        items={field.options}
        itemToStringValue={(opt) => {
          return opt.value;
        }}
        itemToStringLabel={(opt: any) => {
          return opt.label;
        }}
      >
        <ComboboxInput
          className="h-10 w-full min-w-0 rounded-lg border border-input bg-transparent px-2.5 py-1 text-base transition-colors outline-none file:inline-flex file:h-6 file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50 disabled:pointer-events-none disabled:cursor-not-allowed disabled:bg-input/50 disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-3 aria-invalid:ring-destructive/20 md:text-sm dark:bg-input/30 dark:disabled:bg-input/80 dark:aria-invalid:border-destructive/50 dark:aria-invalid:ring-destructive/40"
          placeholder={field.placeholder}
        />
        <ComboboxContent>
          <ComboboxEmpty>No items found</ComboboxEmpty>
          <ComboboxList>
            {(item) => (
              <ComboboxItem key={item?.value} value={item}>
                {item.label}
              </ComboboxItem>
            )}
          </ComboboxList>
        </ComboboxContent>
        {/* <SelectTrigger */}
        {/*   className={cn( */}
        {/*     field.classNames?.select, */}
        {/*     error && "border-destructive", */}
        {/*   )} */}
        {/* > */}
        {/*   <SelectValue /> */}
        {/* </SelectTrigger> */}
        {/**/}
        {/* <SelectContent> */}
        {/*   {field.options?.map((option) => ( */}
        {/*     <SelectItem */}
        {/*       key={String(option.value)} */}
        {/*       value={String(option.value)} */}
        {/*       disabled={option.disabled} */}
        {/*     > */}
        {/*       {option.label} */}
        {/*     </SelectItem> */}
        {/*   ))} */}
        {/* </SelectContent> */}
      </Combobox>
    </DynamicFieldWrapper>
  );
}
