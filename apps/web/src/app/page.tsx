import Link from "next/link";

export default function HomePage() {
  return (
    <main className="mx-auto min-h-screen max-w-4xl p-8">
      <h1 className="text-2xl font-bold">MVP Web Admin</h1>
      <p className="mt-2 text-slate-600">Starter berjalan. Lanjut ke auth dan RBAC implementation.</p>
      <div className="mt-6">
        <Link className="text-brand-700 underline" href="/login">
          Go to Login
        </Link>
      </div>
    </main>
  );
}
