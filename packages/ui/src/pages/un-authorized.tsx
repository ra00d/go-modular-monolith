import { Button } from "@/components/ui/button";
// TODO: add localization
export function UnauthorizedPage() {
  return (
    <main className="error-page">
      <div className="error-shell error-shell-wide">
        <div className="lock-orbit" aria-hidden="true">
          <span>!</span>
        </div>
        <p className="eyebrow">Restricted area</p>
        <h1>You&apos;re not cleared for this view.</h1>
        <p className="error-copy">
          This space is reserved for approved members. Sign in with an account
          that has access, or ask your workspace administrator for an
          invitation.
        </p>
        <div className="error-actions">
          <Button size="lg">
            <a href="/">
              Sign in to continue <span aria-hidden="true">→</span>
            </a>
          </Button>
          <Button variant="outline" size="lg">
            <a href="/">Back to home</a>
          </Button>
        </div>
        <p className="error-hint">
          Error code <strong>ACCESS_REQUIRED</strong>
        </p>
      </div>
    </main>
  );
}
