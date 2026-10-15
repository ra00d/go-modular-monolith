import { PageLayout } from "../page-layout";
import type { FieldValues } from "react-hook-form";
import { DynamicForm, type DynamicFormProps } from "#/components/forms";
import { Button } from "#/components/ui";
import type { ReactNode } from "react";
export interface FormPageProps {
  title?: ReactNode;
  description?: ReactNode;
  actions?: ReactNode[];
  submitButton?: ReactNode;
  submitLable?: ReactNode;
  isEdit?: boolean;
  footer?: ReactNode;
}
export const FormPage = <TFieldValues extends FieldValues = FieldValues>({
  form,
  title,
  description,
  actions,
  isEdit,
  footer,
  submitButton,
  submitLable,
}: FormPageProps & DynamicFormProps<TFieldValues>) => {
  return (
    <PageLayout>
      <PageLayout.Header>
        <PageLayout.HeaderContent>
          {title ? (
            typeof title === "string" ? (
              <PageLayout.Title>{title}</PageLayout.Title>
            ) : (
              title
            )
          ) : (
            <></>
          )}
          {description ? (
            typeof description === "string" ? (
              <PageLayout.Description>{description}</PageLayout.Description>
            ) : (
              description
            )
          ) : (
            <></>
          )}
        </PageLayout.HeaderContent>
        {actions && actions.length > 0 && (
          <PageLayout.Actions>{actions.map((a) => a)}</PageLayout.Actions>
        )}
      </PageLayout.Header>

      <PageLayout.Body>
        <DynamicForm form={form}>
          <div className="flex justify-end">
            {submitButton
              ? submitButton
              : submitLable && <Button type="submit">{submitLable}</Button>}
          </div>
        </DynamicForm>
      </PageLayout.Body>
    </PageLayout>
  );
};
