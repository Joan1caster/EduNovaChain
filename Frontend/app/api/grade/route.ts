const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET() {
  const response = await fetch(`${url}/api/v1/grade`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  return Response.json(data);
}
