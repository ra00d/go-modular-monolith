import { m } from "#/paraglide/messages";
import { createFileRoute } from "@tanstack/react-router";
import { toast } from "@workspace/ui/components";
import { useDynamicForm, type FormField } from "@workspace/ui/components/forms";
import { FormPage } from "@workspace/ui/layouts/pages/form-page";
import z from "zod";

export const Route = createFileRoute("/(app)/form")({
  component: RouteComponent,
});

const fields: FormField[] = [
  {
    name: "name",
    type: "text",
    label: m.label_name(),
    desc: m.customer_fullname(),
    colsSpan: 6,
    classNames: {},
    //@ts-ignore
    validation: z.string(m.validation_required()).min(
      2,
      m.validation_min({
        min: 2,
      }),
    ),
  },

  {
    name: "email",
    type: "email",
    label: m.label_email(),
    colsSpan: 6,
    //@ts-ignore
    validation: z.email(m.validation_email()),
  },

  {
    name: "phone",
    type: "text",
    label: m.label_phone(),
    colsSpan: 6,
    //@ts-ignore
    validation: z
      .string(m.validation_required())
      .min(10, m.validation_min({ min: 10 })),
  },

  {
    name: "type",
    type: "select",
    label: m.customer_type(),
    colsSpan: 6,
    options: [
      {
        label: m.individual(),
        value: "individual",
      },
      {
        label: m.merchant(),
        value: "merchant",
      },
    ],
    //@ts-ignore
    validation: z.enum(
      ["individual", "merchant"],
      m.validation_in({
        opts: [m.individual(), m.merchant()].join("-"),
      }),
    ),
  },

  {
    name: "description",
    type: "textarea",
    label: m.label_description(),
    colsSpan: 12,
    //@ts-ignore
    validation: z.string().optional(),
  },

  {
    name: "active",
    type: "checkbox",
    label: m.label_active(),
    defaultValue: true,
    colsSpan: 12,
    //@ts-ignore
    validation: z.boolean().default(false),
  },
];
function RouteComponent() {
  const handleSubmit = (values: any) => {
    console.log(values);
    toast.add({
      title: "Form submitted",
      description: <pre>{JSON.stringify(values, null, 2)}</pre>,
    });
  };

  const form = useDynamicForm({
    fields,
    onSubmit: handleSubmit,
    defaultValues: {
      name: "",
      email: "",
      type: "individual",
      active: true,
    },
  });

  return (
    <FormPage
      title={m.nav_products()}
      descriprion="create new product"
      form={form}
      submitLable={"Save"}
    />
  );
}
