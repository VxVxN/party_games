import {useState} from "react";
import {AnimatePresence, motion, useMotionValue, useTransform} from "framer-motion";
import './styles.scss'
import CounterCards from "../Counter";

const cards = [
    "Я никогда не использовал поддельное удостоверение личности.",
    "Меня никогда не арестовывали.",
    "Я никогда не унижал себя на свидании.",
    "Я никогда не роняла телефон в унитаз.",
    "Я никогда не касался червя.",
    "Я никогда не был в магазине для взрослых.",
    "Я никогда не флиртовал с кем-то, чтобы получить бесплатную выпивку.",
    "Меня никогда не вырвало на незнакомца, когда он был пьян,",
    "Я никогда не мочился в постель старше 15 лет.",
    "У меня никогда не было сахарного папочки/мамы.",
    "Я никогда не водил машину голым.",
    "Я никогда не бросал пить более двух раз.",
    "Я никогда не бросал курить более двух раз.",
    "Я никогда не плавал голышом в чужом бассейне.",
    "Я никогда не выходил на улицу без одежды.",
    "Я никогда не платил за контент для взрослых.",
    "Я никогда не звонил своим родителям через задницу.",
    "Я никогда не танцевал на столе.",
    "Я никогда не ходил на работу с похмелья.",
    "Я никогда не флиртовала с учителем.",
    "Я никогда не целовался в самолете.",
    "Я никогда не был в стриптиз-клубе.",
];

function Questions() {
    const [index, setIndex] = useState(0);
    const [direction, setDirection] = useState(0);

    const x = useMotionValue(0)
    const xInput = [-100, 0, 100]
    const scale = useTransform(x, [-200, 0, 200], [0.8, 1, 0.8]);
    const background = useTransform(x, xInput, [
        "linear-gradient(180deg, rgba(255,226,242,1) 0%, rgba(151,255,212,1) 100%)",
        "linear-gradient(180deg, #fff 0%, rgb(255, 255, 255) 100%)",
        "linear-gradient(180deg, rgba(255,226,242,1) 0%, rgba(255,151,189,1) 100%)",
    ])

    const paginate = (newDirection: number) => {
        setDirection(newDirection);
        setIndex((prev) => Math.min(Math.max(prev + newDirection, 0), cards.length - 1));
    };

    return <div>
        <CounterCards count={cards.length} current={index}/>
        <div className="questions">
            <AnimatePresence>
                <motion.div
                    key={index}
                    className="card-question absolute w-80 h-40 flex justify-center items-center bg-white rounded-lg shadow-lg text-xl text-slate-950"
                    drag="x"
                    dragConstraints={{left: 0, right: 0}}
                    dragElastic={0.5}
                    style={{x, background, scale}}
                    initial={{x: direction > 0 ? 300 : -300, opacity: 0, scale: 1}}
                    animate={{x: 0, opacity: 1, rotateY: 0}}
                    exit={{opacity: 0}}
                    transition={{type: "spring", stiffness: 500, damping: 30}}
                    whileTap={{scale: 0.95}}
                    onDragEnd={(_, info) => {
                        if (info.offset.x < -100) {
                            // Свайп влево
                            paginate(1);
                        } else if (info.offset.x > 100) {
                            // Свайп вправо
                            paginate(-1);
                        } else {
                            x.set(0, true)
                        }
                    }}
                >
                    {cards[index]}
                </motion.div>
            </AnimatePresence>
        </div>
    </div>
}

export default Questions;