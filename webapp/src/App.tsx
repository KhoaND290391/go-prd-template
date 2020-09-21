import React, { useState, useEffect } from 'react';

import './App.css';
import { Snackbar } from '@material-ui/core';
import Alert from '@material-ui/lab/Alert/Alert';
type PrimeNum = string | number;

type AppState = {
    inputNumber: number,
    calculatedNumber: number,
    toastShow: boolean,
    toastMessage: string,
};

type AppProps = {
    currentInput?: PrimeNum;
    result?: PrimeNum;
};


const App: React.FunctionComponent<AppProps> = (props) => {
    const [state, setState] = useState<AppState>({
        calculatedNumber: 0,
        inputNumber: 0,
        toastMessage: "",
        toastShow: false
    });

    const showToast = (message: string) => {
        setState({ ...state, toastMessage: message, toastShow: true});
    };
    const handleClose = () => { setState((prev) => ({ ...prev, toastMessage: "", toastShow: false}))};
    function raiseError(message: string) {
        showToast(message);
    }
    function handleOnChange(e: React.ChangeEvent<HTMLInputElement>): void {
        const text = e.currentTarget.value;
        try {
            var newValue = Number.parseInt(text, 10);
            if  (isNaN(newValue) || newValue >= Number.MAX_SAFE_INTEGER ) {
                e.preventDefault();
                showToast("Invalid input, range 3-> 2^53 -1");
                return;
            }
            setState((prev) => ({ ...prev, inputNumber: newValue }))
        } catch (ex) {
            e.preventDefault();
            raiseError(ex);
        }
    };
   


    console.log("render, toast: ", state.toastMessage, "---show: ", state.toastShow);

    return (
        <div className="App">
            <header className="App-header">
                <input type="number"
                    max={Number.MAX_SAFE_INTEGER}
                    className="number-field"
                    value={state.inputNumber}
                    onChange={handleOnChange}
                    defaultValue={0}>
                </input>
                <p>Your input = {state.inputNumber}</p>
            </header>

            <Snackbar
                anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
                open={state.toastShow}
                onClose={handleClose}
                autoHideDuration={3000}
                key={"bottom-right"}
            >
                <Alert variant="filled" severity="warning" id={"alert-bottom-right"}>
                    {state.toastMessage ?? "Alert Message Default"}
                </Alert>
            </Snackbar>
        </div>
    );
}
export default App;
