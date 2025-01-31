import {motion} from "framer-motion";
import './styles.scss'

interface PageHeaderProps {
    title: string
    size?: 'small' | 'default'
}

function PageHeader(props: PageHeaderProps) {
    const {title, size} = props;

    return <motion.h1 className='page-header mb-5' initial={{scale: 0, fontSize: `${size === 'small' ? 1.5 : 3}rem`}} animate={{scale: 1}}>{title}</motion.h1>
}

export default PageHeader;