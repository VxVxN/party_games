import { motion } from "framer-motion";
import './styles.scss'
import {usePrev} from "@/hooks/usePrev.ts";
import {memo} from "react";

interface CounterCardPros {
    count: number
    current: number
}

function CounterCards(props: CounterCardPros) {
   const {count, current} = props;
   const prevCurrent = usePrev(current)

    const isNext = prevCurrent !== undefined && current > prevCurrent;

    return <div className='count-display flex justify-center'>
        <motion.div
            key={current} // Когда счетчик изменяется, ключ меняется и компонент перерисовывается
            initial={{y: isNext ? -50 : 50, opacity: 0}} // Начинаем с небольшого смещения вверх
            animate={{y: 0, opacity: 1}} // Плавно переходим в обычную позицию
            exit={{y: isNext ? 50 : -50, opacity: 0}} // Уходим вниз с плавным исчезновением
            transition={{type: "spring", stiffness: 300, damping: 30}} // Настройка плавности
        >
            {current + 1}
        </motion.div>
        <div>/</div>
        <div>{count}</div>
    </div>
}

export default memo(CounterCards);
