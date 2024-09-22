import { cookies } from "next/headers";
import { NextRequest } from "next/server";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "http://172.26.45.135:4455";

export async function POST(request: NextRequest) {
  const body = await request.json();
  const cookie = cookies();
  const Authorization = cookie.get("Authorization")?.value ?? "";

  request.headers.set("Authorization", Authorization);
  const response = await fetch(`${url}/api/v1/ipfs/upload`, {
    method: "POST",
    body: JSON.stringify(body),
    headers: {
      Authorization: Authorization,
    },
  });

  const data = await response.json();

  return Response.json({ data: data.data });
}

export async function GET(request: NextRequest) {
  const hash = request.nextUrl.searchParams.get("hash");
  try {
    const response = await fetch(`${url}/api/v1/ipfs/data/${hash}`);
    const data = await response.json();
    return Response.json(data.data);
  } catch {
    return Response.json({ title: "", content: "", summary: "" });
  }
}
