import z from "zod";

export const loginSchema = z.object({
  email: z.email("invalid_email"),
  password: z.string().min(1, "field_is_required"),
});

export type LoginFormData = z.infer<typeof loginSchema>;
