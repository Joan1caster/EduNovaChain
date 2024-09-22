import { hash } from "crypto";
import { cookies } from "next/headers";
import { NextRequest } from "next/server";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "http://172.26.45.135:4455";

export async function GET(request: NextRequest) {
  const type = request.nextUrl.searchParams.get("type");
  const value = request.nextUrl.searchParams.get("value");
  const cookie = cookies();
  const Authorization = cookie.get("Authorization")?.value ?? "";

  const isCheckStatus = type === "status";
  const fetchUrl = isCheckStatus
    ? `${url}/api/v1/orders/status/${value}`
    : `${url}/api/v1/orders/${value}`;
  const response = await fetch(fetchUrl, {
    headers: {
      Authorization,
    },
  });

  const data = await response.json();

  return Response.json(data);
}

export async function POST(request: Request) {
  const body = await request.json();
  const cookie = cookies();
  const Authorization = cookie.get("Authorization")?.value ?? "";

  const response = await fetch(`${url}/api/v1/orders/buy`, {
    method: "POST",
    cache: "no-store",
    headers: {
      Authorization,
    },
    body: JSON.stringify(body),
  });
  const data = await response.json();

  const statusResonse = await fetch(
    `${url}/api/v1/orders/status/${body.tx_hash}`,
    {
      headers: {
        Authorization,
      },
    }
  );
  const statusJson = await statusResonse.json();
  console.log(statusJson);

  return Response.json(data);
}
