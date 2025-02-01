import { useState } from "react";
import {
  AnimatePresence,
  motion,
  useMotionValue,
  useTransform,
} from "framer-motion";
import "./styles.scss";
import CounterCards from "../Counter";
// import {useGetTopicRecords} from "@/api/getTopicRecords.ts";
// import {RequestTopicRecords} from "@/api/types.ts";
// import {usePrev} from "@/hooks/usePrev.ts";

const cards = [
  "Я никогда не использовал поддельное удостоверение личности.",
  "Меня никогда не арестовывали.",
  "Я никогда не унижал себя на свидании.",
];

interface QuestionsProps {
  topics: string[];
}

function Questions(props: QuestionsProps) {
  const { topics } = props;
  const [index, setIndex] = useState(0);
  const [direction, setDirection] = useState(0);
  const [page, setPage] = useState(1);
  // const prevPage = usePrev(page);

  // const {data: questions, isSuccess, isError, mutate} = useGetTopicRecords({topics, page})

  console.log(page, topics);

  const x = useMotionValue(0);
  const xInput = [-100, 0, 100];
  const scale = useTransform(x, [-200, 0, 200], [0.8, 1, 0.8]);
  const background = useTransform(x, xInput, [
    "linear-gradient(180deg, rgba(255,226,242,1) 0%, rgba(151,255,212,1) 100%)",
    "linear-gradient(180deg, #fff 0%, rgb(255, 255, 255) 100%)",
    "linear-gradient(180deg, rgba(255,226,242,1) 0%, rgba(255,151,189,1) 100%)",
  ]);

  // const requestArgs = useMemo<RequestTopicRecords>(() => {
  //     const args = {
  //         topics,
  //         page: page < 1 ? 1 : page,
  //     }
  //
  //     if (isError && prevPage) {
  //         args.page = prevPage;
  //     }
  //
  //     return args;
  // }, [topics, page, prevPage, isError])

  // useEffect(() => {
  //     // mutate(requestArgs)
  // }, [requestArgs])

  const paginate = (newDirection: number) => {
    setDirection(newDirection);
    setIndex((prev) =>
      Math.min(Math.max(prev + newDirection, 0), cards.length - 1),
    );
  };

  console.log();

  return (
    <div>
      <CounterCards count={cards.length} current={index} />
      <div className="questions">
        <AnimatePresence>
          <motion.div
            key={index}
            className="card-question absolute w-80 h-40 flex justify-center items-center bg-white rounded-lg shadow-lg text-xl text-slate-950"
            drag="x"
            dragConstraints={{ left: 0, right: 0 }}
            dragElastic={0.5}
            style={{ x, background, scale }}
            initial={{ x: direction > 0 ? 300 : -300, opacity: 0, scale: 1 }}
            animate={{ x: 0, opacity: 1, rotateY: 0 }}
            exit={{ opacity: 0 }}
            transition={{ type: "spring", stiffness: 500, damping: 30 }}
            whileTap={{ scale: 0.95 }}
            onDragEnd={(_, info) => {
              if (info.offset.x < -100) {
                if (index === cards.length - 1) {
                  setPage((prevState) => prevState + 1);
                  console.log("fetch next page");
                }

                // Свайп влево
                paginate(1);
              } else if (info.offset.x > 100) {
                if (index === 0) {
                  setPage((prevState) => prevState - 1);
                  console.log("fetch prev page");
                }

                // Свайп вправо
                paginate(-1);
              } else {
                x.set(0, true);
              }
            }}
          >
            {cards[index]}
          </motion.div>
        </AnimatePresence>
      </div>
    </div>
  );
}

export default Questions;
