"use client";

import { useFormStatus } from "react-dom";

export default function SubmitButton({
  children,
}: {
  children: React.ReactNode;
}) {
  const { pending } = useFormStatus();

  return (
    <button
      type="submit"
      aria-disabled={pending}
      className="w-full rounded-md bg-primary px-10 py-2 mx-auto font-semibold text-white shadow-sm hover:bg-primary focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-light/50 disabled:bg-primary-light/50 disabled:cursor-not-allowed"
    >
      {children}
    </button>
  );
}
