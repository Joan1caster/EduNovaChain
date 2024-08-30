'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useEffect, useState } from 'react';

export default function LoginButton({ email }: { email: string }) {
  const pathname = usePathname();
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    setVisible(!/^\/account\/./.test(pathname));
  }, [pathname]);

  if (visible)
    return (
      <button type="submit" className="rounded-md px-3 py-1.5 text-sm">
        {email ? <div>{email}</div> : <Link href="/account/login">Log in</Link>}
      </button>
    );
  else return <></>;
}
