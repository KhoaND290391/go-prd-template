import React, { useState, useEffect } from "react";

import "./App.css";
import { Snackbar } from "@material-ui/core";
import Alert from "@material-ui/lab/Alert/Alert";
import Button from "@material-ui/core/Button";

type PrimeNum = string | number;

type AppState = {
  inputNumber: number;
  calculatedNumber: number;
  toastShow: boolean;
  toastMessage: string;
  validInput: boolean;
  sending: boolean;
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
    toastShow: false,
    validInput: false,
    sending: false,
  });

  useEffect(() => {
    getPrime();
  }, []);
  const getPrime = async () => {
      if (process && process.env) {
        const url = process?.env?.SERVER_API ?? "" + process?.env?.GET_PRIME_URL ?? "";
        const response = await fetch(url, {
            method: 'POST',
            mode: 'no-cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({number: state.inputNumber})
        });
        const data = await response.json();
      }
  }


  const handleClose = () => {
    setState((prev) => ({ ...prev, toastMessage: "", toastShow: false }));
  };

  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    const text = e.currentTarget.value;
    try {
      var newValue = Number.parseInt(text, 10);
      if (isNaN(newValue) || newValue < 1  || newValue >= Number.MAX_SAFE_INTEGER) {
        setState((prev) => ({ ...prev, toastMessage: "Invalid input, range 3 -> 2^53 -1", toastShow: true, validInput: false, }));
        return;
      }
      setState((prev) => ({
        ...prev,
        inputNumber: newValue,
        validInput: true,
      }));
    } catch (ex) {
      setState((prev) => ({ ...prev, validInput: false, toastMessage: ex, toastShow: true  }));
    }
  }

  console.log(
    "render, toast: ",
    state.toastMessage,
    "---show: ",
    state.toastShow
  );

  return (
    <div className="App">
      <header className="App-header">
        <input
          type="number"
          max={Number.MAX_SAFE_INTEGER}
          className="number-field"
          value={state.inputNumber}
          onChange={handleOnChange}
          defaultValue={0}
        ></input>
        <Button variant="contained" disabled={!state.validInput}>
          Find highest lower prime number
        </Button>
        <p>{state.calculatedNumber ? `Result: ${state.calculatedNumber}` : "Waiting you!"}</p>
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
};
export default App;
