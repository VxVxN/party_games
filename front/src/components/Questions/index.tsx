import { useCallback, useEffect, useMemo, useState } from "react";
import {
  AnimatePresence,
  motion,
  PanInfo,
  useMotionValue,
  useTransform,
} from "framer-motion";
import "./styles.scss";
import CounterCards from "../Counter";
import { useGetTopicRecords } from "@/api/getTopicRecords.ts";
import { RequestTopicRecords } from "@/api/types.ts";
import { usePrev } from "@/hooks/usePrev.ts";
import Loader from "@/components/Loader";

interface QuestionsProps {
  topics: string[];
}

function Questions(props: QuestionsProps) {
  const { topics } = props;
  const [index, setIndex] = useState(0);
  const [direction, setDirection] = useState(0);
  const [page, setPage] = useState(1);
  const prevPage = usePrev(page);

  const { data: questions, isSuccess, isError, mutate } = useGetTopicRecords();

  const x = useMotionValue(0);
  const xInput = [-100, 0, 100];
  const scale = useTransform(x, [-200, 0, 200], [0.8, 1, 0.8]);
  const background = useTransform(x, xInput, [
    "linear-gradient(180deg, rgba(255,226,242,1) 0%, rgba(151,255,212,1) 100%)",
    "linear-gradient(180deg, #fff 0%, rgb(255, 255, 255) 100%)",
    "linear-gradient(180deg, rgba(255,226,242,1) 0%, rgba(255,151,189,1) 100%)",
  ]);

  const hasNextPage = useMemo(
    () => questions?.count_page && page < questions.count_page,
    [questions?.count_page, page],
  );

  const requestArgs = useMemo<RequestTopicRecords>(() => {
    return {
      topics,
      page: isError && prevPage ? prevPage : Math.max(page, 1),
    };
  }, [topics, page, prevPage, isError]);
  const prevRequestArgs = usePrev(requestArgs);

  useEffect(() => {
    if (JSON.stringify(requestArgs) !== JSON.stringify(prevRequestArgs)) {
      if (questions?.count_page && page > questions?.count_page) return;

      mutate(requestArgs, {
        onSuccess: (resp) => {
          if (prevPage && prevPage > page) {
            setIndex(resp.records.length ? resp.records.length - 1 : 0);
            setDirection(0);
            return;
          }

          setIndex(0);
          setDirection(0);
        },
      });
    }
  }, [
    mutate,
    page,
    prevPage,
    prevRequestArgs,
    questions?.count_page,
    requestArgs,
  ]);

  const paginate = useCallback(
    (newDirection: number) => {
      setDirection(newDirection);
      setIndex((prev) =>
        Math.min(
          Math.max(prev + newDirection, 0),
          (questions?.records.length || 0) - 1,
        ),
      );
    },
    [questions?.records.length],
  );

  const handleDragEnd = useCallback(
    (_: unknown, info: PanInfo) => {
      if (info.offset.x < -100) {
        if (index === (questions?.records.length || 0) - 1) {
          if (hasNextPage) setPage((prev) => prev + 1);

          return;
        }

        // Свайп влево
        paginate(1);
      } else if (info.offset.x > 100) {
        if (index === 0 && page > 1) {
          setPage((prev) => Math.max(prev - 1, 1));

          return;
        }

        // Свайп вправо
        paginate(-1);
      } else {
        x.set(0, true);
      }
    },
    [index, questions?.records.length, paginate, hasNextPage, page, x],
  );

  return (
    <div>
      <CounterCards count={questions?.records.length || 0} current={index} />
      <div className="flex gap-2 w-full justify-center text-neutral-400">
        <span>Этап</span>
        <span>{page}</span>
      </div>
      <div className="questions">
        {isSuccess ? (
          questions?.records.length === 0 ? (
            <div>Empty</div>
          ) : (
            <AnimatePresence>
              <motion.div
                key={index}
                className="card-question absolute w-80 h-40 flex justify-center items-center bg-white rounded-lg shadow-lg text-xl text-slate-950"
                drag="x"
                dragConstraints={{ left: 0, right: 0 }}
                dragElastic={0.5}
                style={{ x, background, scale }}
                initial={{
                  x: direction > 0 ? 300 : -300,
                  opacity: 0,
                  scale: 1,
                }}
                animate={{ x: 0, opacity: 1, rotateY: 0 }}
                exit={{ opacity: 0 }}
                transition={{ type: "spring", stiffness: 500, damping: 30 }}
                whileTap={{ scale: 0.95 }}
                onDragEnd={handleDragEnd}
              >
                {questions?.records[index]}
              </motion.div>
            </AnimatePresence>
          )
        ) : (
          <Loader />
        )}
      </div>
    </div>
  );
}

export default Questions;
