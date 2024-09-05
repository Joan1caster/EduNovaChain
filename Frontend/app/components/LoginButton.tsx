"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

export default function LoginButton({ email }: { email: string }) {
  const pathname = usePathname();
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    setVisible(!/^\/account\/./.test(pathname));
  }, [pathname]);

  if (visible)
    return (
      <button className="rounded-md bg-primary px-10 py-2 mx-auto font-semibold text-white shadow-sm hover:bg-primary focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-light/50 disabled:bg-primary-light/50 disabled:cursor-not-allowed">
        {email ? <div>{email}</div> : <Link href="/account/login">登录</Link>}
      </button>
    );
  else return <></>;
}
