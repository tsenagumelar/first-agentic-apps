import { InputHTMLAttributes } from "react";

type Props = InputHTMLAttributes<HTMLInputElement> & {
  label: string;
  error?: string;
};

export function FormField({ label, error, ...props }: Props) {
  return (
    <label className="block space-y-1">
      <span className="text-sm font-medium text-slate-700">{label}</span>
      <input className="w-full rounded-lg border border-slate-300 px-3 py-2" {...props} />
      {error ? <span className="text-xs text-red-600">{error}</span> : null}
    </label>
  );
}
