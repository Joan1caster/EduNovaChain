"use client";
import Image from "next/image";
import { TagType } from "@/app/types";
import { useState, useEffect, useReducer } from "react";

type Props = {
  followTopics: TagType[];
  onUpdateFollowTopics: (topics: TagType[]) => void;
};

const grades: TagType[] = [
  { name: "全部", key: -1 },
  { name: "小学", key: 0 },
  { name: "初中", key: 1 },
  { name: "高中", key: 2 },
  { name: "高校", key: 3 },
];

const subjects: TagType[] = [
  { name: "哲学", key: 0 },
  { name: "法学", key: 1 },
  { name: "经济学", key: 2 },
  { name: "文学", key: 3 },
  { name: "理学", key: 4 },
  { name: "教育学", key: 5 },
  { name: "历史学", key: 6 },
  { name: "工学", key: 7 },
  { name: "农学", key: 5 },
  { name: "医学", key: 6 },
  { name: "军事学", key: 7 },
];

const topics: TagType[] = [
  { name: "电子", key: 0 },
  { name: "航天", key: 1 },
  { name: "设计", key: 2 },
  { name: "机器学习", key: 3 },
  { name: "市场营销", key: 4 },
  { name: "古代文学", key: 5 },
  { name: "现代文学", key: 6 },
  { name: "中国近代史", key: 7 },
  { name: "美国近代史", key: 8 },
  { name: "植物学", key: 9 },
  { name: "军事学", key: 10 },
];

type InitialState = {
  topicKeys: number[];
  topics: TagType[];
};
type Action =
  | { type: "init"; payload: TagType[] }
  | { type: "add"; payload: TagType }
  | { type: "remove"; payload: TagType };

const initialState: InitialState = {
  topicKeys: [],
  topics: [],
};

function reducer(state: InitialState, action: Action) {
  switch (action.type) {
    case "init": {
      state.topicKeys = action.payload.map((item) => item.key);
      state.topics = action.payload.slice();
      return state;
    }
    case "add": {
      return {
        topicKeys: [...state.topicKeys, action.payload.key],
        topics: [...state.topics, action.payload],
      };
    }
    case "remove": {
      const position = state.topicKeys.indexOf(action.payload.key);
      state.topicKeys.splice(position, 1);
      state.topics.splice(position, 1);
      return Object.assign({}, state);
    }
  }
}

export default function ChangeFollowTopic({
  followTopics,
  onUpdateFollowTopics,
}: Props) {
  const [state, dispatch] = useReducer(reducer, initialState);
  const [currentGrade, setCurrentGrade] = useState(-1);
  const [currentSubject, setCurrentSubject] = useState(0);

  const onChangeGrade = (grade: TagType) => {
    setCurrentGrade(grade.key);
  };
  const onChangeSubject = (subject: TagType) => {
    setCurrentSubject(subject.key);
  };

  useEffect(() => {
    dispatch({ type: "init", payload: followTopics });
  });
  return (
    <div>
      <div className="flex justify-between items-center">
        <div>我关注的主题</div>
        <button
          className="px-4 py-1 bg-blue-400 rounded-sm text-[0.7rem] text-white"
          onClick={() => onUpdateFollowTopics(state.topics)}
        >
          完成
        </button>
      </div>
      <div className="flex my-4">
        {/* choiced topics start */}
        <div className="w-1/3 *:inline-block gap-2 *:text-[0.7rem] *:font-light *:py-0.5 *:px-3 *:mr-2 *:mb-2 *:border *:rounded-full">
          {state.topics.map((item) => (
            <div
              key={item.key}
              className={
                item.key === 0
                  ? "border-blue-200 text-blue-200"
                  : "flex items-center gap-2 border-blue-400 text-blue-400"
              }
              onClick={() =>
                item.key !== 0 && dispatch({ type: "remove", payload: item })
              }
            >
              <span>{item.name}</span>
              {item.key !== 0 && (
                <svg
                  t="1725173545826"
                  className="inline-block mb-0.5 ml-1"
                  viewBox="0 0 1024 1024"
                  version="1.1"
                  xmlns="http://www.w3.org/2000/svg"
                  p-id="1692"
                  width="10"
                  height="10"
                >
                  <path
                    d="M557.12 512l313.6-313.28a32.128 32.128 0 1 0-45.44-45.44L512 466.88l-313.28-313.6a32.128 32.128 0 0 0-45.44 45.44L466.88 512l-313.6 313.28a32 32 0 0 0 0 45.44 32 32 0 0 0 45.44 0L512 557.12l313.28 313.6a32 32 0 0 0 45.44 0 32 32 0 0 0 0-45.44L557.12 512z"
                    fill="#60a5fa"
                    p-id="1693"
                  ></path>
                </svg>
              )}
            </div>
          ))}
        </div>
        {/* choiced topics end */}

        {/* all topics start */}
        <div className="pl-4 border-l border-l-gray-100">
          {/* grade start */}
          <div>
            <span className="text-sm pr-4">年级</span>
            {grades.map((item) => (
              <span
                key={item.key}
                className={`${currentGrade === item.key ? "text-blue-500 bg-gray-100" : ""} inline-block py-0.5 px-2 rounded-full text-[0.7rem] font-light text-gray-500 cursor-pointer hover:bg-gray-100`}
                onClick={() => onChangeGrade(item)}
              >
                {item.name}
              </span>
            ))}
          </div>
          {/* grade end */}
          {/* subject start */}

          <div className="my-2">
            <span className="text-sm pr-4">学科</span>
            {subjects.map((item) => (
              <span
                key={item.key}
                className={`${currentSubject === item.key ? "text-blue-500 bg-gray-100" : ""} inline-block py-0.5 px-2 rounded-full text-[0.7rem]  font-light text-gray-500 cursor-pointer hover:bg-gray-100`}
                onClick={() => onChangeSubject(item)}
              >
                {item.name}
              </span>
            ))}
          </div>
          {/* sugject end */}
          {/* topic start */}
          <div className="flex flex-wrap gap-2">
            {topics.map((item) => (
              <>
                {state.topicKeys.includes(item.key) ? (
                  <div
                    key={item.key}
                    className="py-1 px-2 rounded-full text-[0.7rem] font-light bg-blue-400 text-white"
                  >
                    {item.name}
                  </div>
                ) : (
                  <div
                    key={item.key}
                    className="py-1 px-2 rounded-full text-[0.7rem] font-light border border-blue-400 text-blue-400 cursor-pointer"
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
                        fill="#60a5fa"
                        p-id="1842"
                      ></path>
                    </svg>
                    {item.name}
                  </div>
                )}
              </>
            ))}
          </div>
          {/* topic end */}
        </div>
        {/* all topics end */}
      </div>
    </div>
  );
}
