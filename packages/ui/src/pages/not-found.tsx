import { Button, buttonVariants } from "@/components/ui/button";
import { useRouter } from "@tanstack/react-router";
import { Home, ArrowLeft } from "lucide-react";

export function NotFound() {
  const router = useRouter();
  return (
    <main className="flex min-h-screen items-center justify-center px-6 w-full">
      <div className="w-full max-w-md text-center">
        <p className="text-8xl font-bold tracking-tight text-muted-foreground/30">
          404
        </p>

        <h1 className="mt-4 text-3xl font-semibold tracking-tight">
          Page not found
        </h1>

        <p className="mt-3 text-muted-foreground">
          Sorry, we couldn&apos;t find the page you&apos;re looking for. It may
          have been moved or no longer exists.
        </p>

        <div className="mt-8 flex items-center justify-center gap-3">
          <Button variant="outline" onClick={() => router.history.back()}>
            <ArrowLeft />
            Go back
          </Button>

          <a href="/" className={buttonVariants({})}>
            <Home />
            Go home
          </a>
        </div>
      </div>
    </main>
  );
}
