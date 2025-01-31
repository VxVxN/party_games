import { motion } from "framer-motion";
import './styles.scss'

interface CounterCardPros {
    count: number
    current: number
}

function CounterCards(props: CounterCardPros) {
   const {count, current} = props;

    return <div className='count-display flex justify-center'>
        <motion.div
            key={current} // Когда счетчик изменяется, ключ меняется и компонент перерисовывается
            initial={{y: -50, opacity: 0}} // Начинаем с небольшого смещения вверх
            animate={{y: 0, opacity: 1}} // Плавно переходим в обычную позицию
            exit={{y: 50, opacity: 0}} // Уходим вниз с плавным исчезновением
            transition={{type: "spring", stiffness: 300, damping: 30}} // Настройка плавности
        >
            {current + 1}
        </motion.div>
        <div>/</div>
        <div>{count}</div>
    </div>
}

export default CounterCards;
