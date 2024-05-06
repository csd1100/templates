import { createRoot } from 'react-dom/client';
import './index.css';
import Sum from './components/Sum';

const root = createRoot(document.getElementById('root')!);
root.render(
    <div className="container">
        <Sum />
        <div></div>
    </div>
);
