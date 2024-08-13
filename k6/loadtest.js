import http from "k6/http";
import { sleep, check } from "k6";

// export let options = {
//     vus: 100, // Number of virtual users
//     duration: "30s", // Duration of the test
// };

export let options = {
    stages: [
        { duration: "10s", target: 10 },  // Start with 10 VUs
        { duration: "10s", target: 20 },  // Then increase to 20 VUs
        { duration: "10s", target: 40 },  // Then increase to 40 VUs
        { duration: "10s", target: 80 },  // Then increase to 80 VUs
        { duration: "10s", target: 100 }, // Finally, ramp up to 100 VUs
    ],
};


export default function () {
    let res = http.get("http://127.0.0.1:1323/alive");

    check(res, {
        success: (res) => res.status == 200,
        succesString: (res) => res.body.trim() == "yes"
    });

    sleep(1);
}