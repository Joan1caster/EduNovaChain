import { cookies } from "next/headers";
import { NextRequest } from "next/server";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "http://172.26.45.135:4455";

export async function GET(request: NextRequest) {
  const wallet = request.nextUrl.searchParams.get("wallet");
  const cookie = cookies();
  const Authorization = cookie.get("Authorization");
  // if (Authorization) {
  //   return Response.json("");
  // } else {
  const response = await fetch(`${url}/api/v1/siweMessage?wallet=${wallet}`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  return Response.json(data);
  // }
}
