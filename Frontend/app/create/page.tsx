"use client";
import { TagType } from "@/app/types";
import {
  useState,
  useEffect,
  useReducer,
  useRef,
  LegacyRef,
  MutableRefObject,
} from "react";
import { useFormState } from "react-dom";
import Image from "next/image";
import { createNFT } from "./actions";
import SubmitButton from "../components/SubmitButton";
import { Dialog, DialogPanel } from "@headlessui/react";
import { useAsyncEffect } from "ahooks";
import { useWriteContract, useReadContract } from "wagmi";
import { ABIS, ContractConfig } from "../abis";

type Props = {
  followTopics: TagType[];
  onUpdateFollowTopics: (topics: TagType[]) => void;
};

const grades: TagType[] = [
  { name: "小学", id: 0 },
  { name: "初中", id: 1 },
  { name: "高中", id: 2 },
  { name: "高校", id: 3 },
];

const subjects: TagType[] = [
  { name: "哲学", id: 0 },
  { name: "法学", id: 1 },
  { name: "经济学", id: 2 },
  { name: "文学", id: 3 },
  { name: "理学", id: 4 },
  { name: "教育学", id: 5 },
  { name: "历史学", id: 6 },
  { name: "工学", id: 7 },
  { name: "农学", id: 8 },
  { name: "医学", id: 9 },
  { name: "军事学", id: 10 },
];

const topics: TagType[] = [
  { name: "电子", id: 0 },
  { name: "航天", id: 1 },
  { name: "设计", id: 2 },
  { name: "机器学习", id: 3 },
  { name: "市场营销", id: 4 },
  { name: "古代文学", id: 5 },
  { name: "现代文学", id: 6 },
  { name: "中国近代史", id: 7 },
  { name: "美国近代史", id: 8 },
  { name: "植物学", id: 9 },
  { name: "军事学", id: 10 },
];

type InitialState = {
  subjectKeys: number[];
  subjects: TagType[];
};
type Action =
  | { type: "init"; payload: TagType[] }
  | { type: "add"; payload: TagType }
  | { type: "remove"; payload: TagType };

const initialState: InitialState = {
  subjectKeys: [],
  subjects: [],
};

function reducer(state: InitialState, action: Action) {
  switch (action.type) {
    case "init": {
      state.subjectKeys = action.payload.map((item) => item.id);
      state.subjects = action.payload.slice();
      return state;
    }
    case "add": {
      return {
        subjectKeys: [...state.subjectKeys, action.payload.id],
        subjects: [...state.subjects, action.payload],
      };
    }
    case "remove": {
      const position = state.subjectKeys.indexOf(action.payload.id);
      state.subjectKeys.splice(position, 1);
      state.subjects.splice(position, 1);
      return Object.assign({}, state);
    }
  }
}
type InitialFormState = {
  message: string;
  type: "success" | "fail";
  data?: number;
};
const initialFormState: InitialFormState = { message: "", type: "fail" };

export default function Page() {
  const [state, dispatch] = useReducer(reducer, initialState);
  const { data, isError, isPending, writeContractAsync } = useWriteContract();

  const [formState, formAction] = useFormState(createNFT, initialFormState);

  const [isOpen, setIsOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const [currentGrade, setCurrentGrade] = useState<number | null>();
  const [gradeList, setGradeList] = useState<TagType[]>([]);
  const [subjectList, setSubjectList] = useState<TagType[]>([]);

  const {
    data: balance,
    isLoading,
    status,
    error,
  } = useReadContract({
    ...ContractConfig,
    functionName: "balanceOf",
    args: ["0xE460De64beecE8a6f5970931C6bb3277f5Cf5c77" as `0x${string}`],
  });
  useAsyncEffect(async () => {
    console.log(balance, isLoading, status, error);
    const response = await (await fetch("/api/grade")).json();
    if (response.count > 0) {
      setGradeList(response.data);
      setCurrentGrade(response.data[0].id);
    }
  });

  useAsyncEffect(async () => {
    if (currentGrade) {
      const response = await (
        await fetch(`/api/subject?id=${currentGrade}`)
      ).json();
      if (response.count > 0) {
        setSubjectList(response.data);
      }
    }
  }, [currentGrade]);

  const onChangeGrade = (grade: TagType) => {
    setCurrentGrade(grade.id);
  };

  const onClose = () => {
    if (loading) return;
    if (formState.type === "success") setIsOpen(false);
  };

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    // setIsLoading(true);
    // setError(null); // Clear previous errors when a new request starts

    try {
      const formData = new FormData(event.currentTarget);

      const data = {
        title: formData.get("title"),
        summary: formData.get("summary"),
        content: formData.get("content"),
        topic: formData.get("topic"),
        grade: formData.get("grade"),
        subject: formData.get("subject"),
        price: formData.get("price"),
      };
      console.log(data);
      const tokenId = await writeContractAsync({
        ...ContractConfig,
        functionName: "createInnovation",
        args: ["1", "1", "1", BigInt(Number(data.price)) * BigInt(1e18), true],
      });
      // const response = await fetch("/api/submit", {
      //   method: "POST",
      //   body: formData,
      // });

      // if (!response.ok) {
      //   throw new Error("Failed to submit the data. Please try again.");
      // }

      // // Handle response if necessary
      // const data = await response.json();
      // ...
    } catch (error) {
      // Capture the error message to display to the user
      // setError(error.message);
      console.error(error);
    } finally {
      // setIsLoading(false);
    }
  }

  return (
    <>
      <form className="mx-auto" onSubmit={onSubmit}>
        <div className="space-y-12">
          <div>
            <h2 className="text-base font-semibold leading-7 text-center text-[#666]">
              发布创意{balance}
              {isLoading}
            </h2>

            <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
              <div className="col-span-full">
                <label className="block leading-6 text-[#666]">标题</label>
                <div className="mt-2">
                  <input
                    id="title"
                    name="title"
                    required
                    placeholder="请输入创意名称，仅支持中文、英文、数字字符"
                    className="block w-full rounded-md border-0 py-1.5 text-[#666] shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-[#ccc] focus:ring-2 focus:ring-inset focus:ring-primary sm:text-sm sm:leading-6"
                  />
                </div>
              </div>
              <div className="col-span-full">
                <label className="block leading-6 text-[#666]">摘要</label>
                <div className="mt-2">
                  <input
                    id="summary"
                    name="summary"
                    required
                    placeholder="请输入摘要内容，明确描述清楚该创意亮点及解决的问题"
                    className="block w-full rounded-md border-0 py-1.5 text-[#666] shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-[#ccc] focus:ring-2 focus:ring-inset focus:ring-primary sm:text-sm sm:leading-6"
                  />
                </div>
              </div>
              <div className="col-span-full">
                <label className="block leading-6 text-[#666]">
                  主题关键词
                </label>
                <div className="mt-2">
                  <input
                    id="topic"
                    name="topic"
                    required
                    placeholder="请输入1-5个关键词，请用英文;隔开"
                    className="block w-full rounded-md border-0 py-1.5 text-[#666] shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-[#ccc] focus:ring-2 focus:ring-inset focus:ring-primary sm:text-sm sm:leading-6"
                  />
                </div>
              </div>
              <div className="col-span-full">
                <div className="flex justify-between">
                  <label className="block leading-6 text-[#666]">
                    创意详设
                  </label>
                  <p className="text-[#666]">
                    <input
                      id="isFree"
                      type="checkbox"
                      className="mr-2 rounded-sm border-[#666] focus:shadow-none focus:ring-offset-0 focus:ring-0"
                    />
                    免费公开
                  </p>
                </div>
                <div className="mt-2">
                  <textarea
                    id="content"
                    name="content"
                    rows={5}
                    required
                    placeholder="请输入创意的详细内容、使用指南和预期效果，字数不限，可作为收益内容。"
                    className="block w-full rounded-md border-0 py-1.5 text-[#666] shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-[#ccc] focus:ring-2 focus:ring-inset focus:ring-primary sm:text-sm sm:leading-6"
                  ></textarea>
                </div>
              </div>
            </div>
          </div>

          <div className="sm:col-span-3">
            <label className="block leading-6 text-[#666]">学科分类</label>
            <div className="mt-2 flex gap-2 items-start">
              <select
                id="grade"
                name="grade"
                autoComplete="grade-name"
                className="block w-20 rounded-md border-0 py-1.5 text-[#666] shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-primary sm:max-w-xs sm:text-sm sm:leading-6"
              >
                {grades.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.name}
                  </option>
                ))}
              </select>

              <div className="flex-1 rounded-md shadow-sm border bg-white py-1 px-2 flex gap-2 flex-wrap items-center">
                {/* <input
                  id="subject"
                  name="subject"
                  value={state.subjectKeys.join(",")}
                  className="hidden"
                /> */}
                {state.subjects.map((item) => (
                  <div
                    key={item.id}
                    className="flex items-center gap-2 text-sm px-4 py-0.5 border rounded-full border-primary bg-primary-light text-primary"
                    onClick={() =>
                      item.id !== 0 &&
                      dispatch({ type: "remove", payload: item })
                    }
                  >
                    <span>{item.name}</span>
                    <svg
                      t="1725173545826"
                      className="inline-block mb-0.5 ml-0.5"
                      viewBox="0 0 1024 1024"
                      version="1.1"
                      xmlns="http://www.w3.org/2000/svg"
                      p-id="1692"
                      width="10"
                      height="10"
                    >
                      <path
                        d="M557.12 512l313.6-313.28a32.128 32.128 0 1 0-45.44-45.44L512 466.88l-313.28-313.6a32.128 32.128 0 0 0-45.44 45.44L466.88 512l-313.6 313.28a32 32 0 0 0 0 45.44 32 32 0 0 0 45.44 0L512 557.12l313.28 313.6a32 32 0 0 0 45.44 0 32 32 0 0 0 0-45.44L557.12 512z"
                        fill="#1474FC"
                        p-id="1693"
                      ></path>
                    </svg>
                  </div>
                ))}
                <input
                  placeholder="支持新增专业分类，多个专业分类用英文;隔开，回车键结束"
                  className="text-sm w-[400px] py-0.5 placeholder:text-[#CCC] border-none outline-none shadow-none focus:shadow-none focus:ring-offset-0 focus:ring-0"
                />
              </div>
            </div>

            {/* subject start */}
            <div className="flex flex-wrap gap-2 mt-4">
              {subjects.map((item) => (
                <div key={item.id}>
                  {state.subjectKeys.includes(item.id) ? (
                    <div
                      key={item.id}
                      className="py-1 px-2 rounded-full text-sm font-light bg-primary text-white"
                    >
                      {item.name}
                    </div>
                  ) : (
                    <div
                      key={item.id}
                      className="py-1 px-2 rounded-full text-sm font-light border border-primary text-primary cursor-pointer"
                      onClick={() => dispatch({ type: "add", payload: item })}
                    >
                      <svg
                        t="1725175223714"
                        className="inline-block mb-0.5 mr-1"
                        viewBox="0 0 1024 1024"
                        version="1.1"
                        xmlns="http://www.w3.org/2000/svg"
                        p-id="1841"
                        width="12"
                        height="12"
                      >
                        <path
                          d="M892.16 480H544V131.84a32 32 0 0 0-64 0V480H131.84a32 32 0 0 0 0 64H480v348.16a32 32 0 1 0 64 0V544h348.16a32 32 0 1 0 0-64z"
                          fill="#1474FC"
                          p-id="1842"
                        ></path>
                      </svg>
                      {item.name}
                    </div>
                  )}
                </div>
              ))}
            </div>
            {/* subject end */}
          </div>
          <div className="sm:col-span-3">
            <label className="block leading-6 text-[#666]">售价</label>
            <div className="w-1/3 mt-2 rounded-md shadow-sm border bg-white py-1 px-2 flex flex-wrap items-center">
              <input
                id="price"
                name="price"
                required
                placeholder="0.00"
                className="text-sm text-right flex-1 py-0.5 placeholder:text-[#CCC] border-none outline-none shadow-none focus:shadow-none focus:ring-offset-0 focus:ring-0"
              />
              <span className="text-[#CCC]">ETH</span>
            </div>
          </div>
        </div>

        <div className="w-[174px] mx-auto mt-10">
          <SubmitButton>发布</SubmitButton>

          <Dialog open={isOpen} onClose={onClose} className="relative z-5">
            <div className="fixed inset-0 flex w-screen h-screen items-center justify-center p-4 bg-black/40">
              <DialogPanel className="w-[496px] py-12 flex flex-col justify-between items-center gap-6 space-y-4 bg-white rounded-lg overflow-hidden">
                {loading && (
                  <>
                    <p className="text-[#333] text-lg">查重校对中，请等待…</p>
                    <div className="flex w-14 h-24 relative space-x-2">
                      <div className="w-2.5 h-2.5 absolute left-2 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite]"></div>

                      <div className="w-2.5 h-2.5 absolute left-4 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite] [animation-delay:0.2s]"></div>

                      <div className="w-2.5 h-2.5 absolute left-8 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite] [animation-delay:0.4s]"></div>

                      <div className="w-2.5 h-2.5 absolute left-12 bottom-0 bg-blue-500 rounded-full animate-[loading_1s_ease-in-out_infinite] [animation-delay:0.6s]"></div>
                    </div>
                  </>
                )}
                {!loading && formState.type === "success" && (
                  <>
                    <p className="text-[#333] text-lg">查重校对已通过</p>
                    <Image
                      src={"/images/slice/check_ok.png"}
                      width={48}
                      height={48}
                      alt="check ok"
                    />
                  </>
                )}
                {!loading && formState.type === "fail" && (
                  <>
                    <p className="text-[#333] text-lg">
                      查重校对结果：
                      <span className="text-[#F06A6A]">
                        {formState.data ?? 0}%
                      </span>
                      ，请返回修改！
                    </p>
                    <Image
                      src={"/images/slice/check_fail.png"}
                      width={48}
                      height={48}
                      alt="check fail"
                    />
                    <div
                      onClick={() => setIsOpen(false)}
                      className="rounded-md bg-primary px-6 py-2 mx-auto font-semibold text-white shadow-sm cursor-pointer hover:bg-primary focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-light/50 "
                    >
                      返回修改
                    </div>
                  </>
                )}
              </DialogPanel>
            </div>
          </Dialog>
        </div>
      </form>
    </>
  );
}
