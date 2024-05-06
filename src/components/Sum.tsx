import { useEffect, useState } from 'react';
const Sum = () => {
    const [a, setA] = useState('0');
    const [b, setB] = useState('0');
    const [sum, setSum] = useState(0);

    async function add(a: string, b: string) {
        if (isNaN(+a) || isNaN(+b)) {
            setSum(0);
            return;
        }
        const sum = await window.api.add(+a, +b);
        setSum(sum);
    }

    useEffect(() => {
        add(a, b);
    }, [a, b]);

    return (
        <div className="container mx-auto my-4 items-center">
            <span className="inline-block">
                <label
                    htmlFor="a"
                    className="my-2 ml-2 h-6 w-12 rounded-lg rounded-r-none border border-current bg-blue-100 px-2 text-blue-500"
                >
                    A
                </label>
                <input
                    type="text"
                    id="a"
                    className={
                        `focus:shadow-outline my-2 mr-2 h-6 w-96 rounded-lg rounded-l-none border bg-blue-100/50 px-4 text-blue-500 outline-none focus:bg-blue-100 focus:ring-2` +
                        (isNaN(+a)
                            ? ' border-red-500 focus:ring-red-200'
                            : ' border-current focus:ring-blue-200')
                    }
                    value={a}
                    onChange={(e) => {
                        setA(e.target.value);
                    }}
                ></input>
            </span>
            +
            <span className="inline-block">
                <label
                    htmlFor="b"
                    className="my-2 ml-2 h-6 w-12 rounded-lg rounded-r-none border border-current bg-blue-100 px-2 text-blue-500"
                >
                    B
                </label>
                <input
                    type="text"
                    id="b"
                    className={
                        `focus:shadow-outline my-2 mr-2 h-6 w-96 rounded-lg rounded-l-none border bg-blue-100/50 px-4 text-blue-500 outline-none focus:bg-blue-100 focus:ring-2` +
                        (isNaN(+b)
                            ? ' border-red-500 focus:ring-red-200'
                            : ' border-current focus:ring-blue-200')
                    }
                    value={b}
                    onChange={(e) => {
                        setB(e.target.value);
                    }}
                ></input>
            </span>
            =
            <input
                type="text"
                id="sum"
                className="focus:shadow-outline m-2 h-6 w-96 rounded-lg border border-current bg-blue-100/50 px-4 text-blue-500 outline-none focus:bg-blue-100 focus:ring-2"
                value={sum}
                readOnly
            ></input>
        </div>
    );
};
export default Sum;
