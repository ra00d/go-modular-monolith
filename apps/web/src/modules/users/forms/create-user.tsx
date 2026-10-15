import { useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import {
	Button,
	DynamicForm,
	type FormField,
	toast,
	useDynamicForm,
} from "@workspace/ui/components";
import { cn } from "@workspace/ui/lib/utils";
import z from "zod";
import { m } from "#/paraglide/messages";
import { $api } from "#/services/api.service";
import type { RequestBody } from "#/types/helpers";

const fields: FormField<RequestBody<"/api/admin/users", "POST">>[] = [
	{
		name: "name",
		label: m.label_name(),
		validation: z.string(m.validation_required()),
	},
	{
		name: "email",
		type: "email",
		label: m.label_email(),
		validation: z
			.string(m.validation_email())
			.email(m.validation_email())
			.or(z.literal("")),
	},
	{
		name: "phone",
		type: "tel",
		label: m.label_phone(),
		validation: z.string().or(z.literal("")),
	},
	{
		name: "username",
		label: m.label_username(),
		validation: z.string().or(z.literal("")),
	},
	{
		name: "password",
		type: "password",
		label: m.label_password(),
		validation: z
			.string(m.validation_required())
			.min(8, m.validation_password({ min: 8 })),
	},
	{
		name: "role",
		type: "number",
		label: m.label_role(),
		defaultValue: 1,
		validation: z.coerce
			.number()
			.int()
			.min(1, m.validation_min({ min: 1 })),
	},
];

export function CreateUserForm({ className }: React.ComponentProps<"form">) {
	const navigate = useNavigate();
	const queryClient = useQueryClient();
	const { mutate, isPending } = $api.useMutation("post", "/api/admin/users", {
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["get", "/api/admin/users"] });
			toast.success({ title: m.success_user_created() });
			navigate({ to: "/users" });
		},
		onError: (error) => {
			console.log(error);
			toast.error({ title: m.error_generic() });
		},
	});
	const form = useDynamicForm({
		defaultValues: {
			name: "",
			email: "",
			phone: "",
			username: "",
			password: "",
			role: 1,
		},
		fields,
		onSubmit: (v) => {
			mutate({
				body: {
					name: v.name,
					email: v.email || undefined,
					phone: v.phone || undefined,
					username: v.username || undefined,
					password: v.password,
					role: Number(v.role),
				},
			});
		},
	});
	return (
		<div className="flex flex-col gap-3">
			<div className="flex flex-col items-center gap-1 text-center">
				<h1 className="text-2xl font-bold">Create a user</h1>
				<p className="text-sm text-balance text-muted-foreground">
					Fill in the details to create a new user account
				</p>
			</div>

			<DynamicForm form={form} className={cn("gap-3 flex flex-col", className)}>
				<Button type="submit" className="w-full" loading={isPending}>
					{m.action_create()}
				</Button>
			</DynamicForm>
		</div>
	);
}
