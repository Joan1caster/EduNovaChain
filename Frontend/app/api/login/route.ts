import { cookies } from "next/headers";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function POST(request: Request) {
  const body = await request.json();
  const response = await fetch(`${url}/api/v1/login`, {
    method: "POST",
    cache: "no-store",
    headers: {
      "ngrok-skip-browser-warning": "off",
    },
    body: JSON.stringify(body),
  });
  const data = await response.json();
  const cookie = cookies();
  cookie.set("Authorization", data.data.token);

  return Response.json(data);
}
