import { m } from "#/paraglide/messages";
import { $api } from "#/services/api.service";
import type { RequestBody } from "#/types/helpers";
import { useNavigate } from "@tanstack/react-router";
import {
  Button,
  DynamicForm,
  toast,
  useDynamicForm,
  type FormField,
} from "@workspace/ui/components";
import { useAuthStore } from "@workspace/ui/stores/auth-store";
import z from "zod";
const fields: FormField<RequestBody<"/api/auth/sign-in", "POST">>[] = [
  {
    name: "identifier",
    label: m.label_username(),
    validation: z.string(m.validation_required()),
  },
  {
    name: "password",
    type: "password",
    label: m.label_password(),
    validation: z
      .string(m.validation_required())
      .min(8, m.validation_password({ min: 8 })),
  },
];
export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"form">) {
  const navigate = useNavigate();
  const { setUser } = useAuthStore();
  const { mutate, isPending } = $api.useMutation("post", "/api/auth/sign-in", {
    onSuccess: (data: any) => {
      if (data.data)
        setUser({
          id: data.data?.id,
          name: data.data?.name,
          email: data.data?.email,
          premissions: data.data?.permissions,
          roles: data.data?.roles,
        });
      navigate({
        to: "/form",
      });
    },
    onError: (error) => {
      console.log(error);
      toast.error({
        title: m.error_generic(),
      });
    },
  });
  const form = useDynamicForm({
    defaultValues: {
      password: "",
      identifier: "",
    },
    fields,
    onSubmit: (v) => {
      mutate({
        body: {
          identifier: v.identifier,
          password: v.password,
        },
      });
    },
  });
  const p = form.watch("password");
  return (
    <div className="flex flex-col gap-3">
      <div className="flex flex-col items-center gap-1 text-center">
        <h1 className="text-2xl font-bold">Login to your account</h1>
        <p className="text-sm text-balance text-muted-foreground">
          Enter your email below to login to your account
        </p>
      </div>

      <DynamicForm form={form} className="gap-3 flex flex-col">
        <Button type="submit" className="w-full" loading={isPending}>
          Submit
        </Button>
      </DynamicForm>
    </div>
  );
}
