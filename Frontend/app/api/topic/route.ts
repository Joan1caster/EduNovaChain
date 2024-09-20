import { NextRequest } from "next/server";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET(request: NextRequest) {
  const gradeId = request.nextUrl.searchParams.get("gradeId");
  const subjectId = request.nextUrl.searchParams.get("subjectId");
  const response = await fetch(`${url}/api/v1/topic/subjectAndGrade`, {
    method: "POST",
    cache: "no-store",
    body: JSON.stringify({
      gradeId,
      subjectId,
    }),
  });
  const data = await response.json();
  return Response.json(data);
}
