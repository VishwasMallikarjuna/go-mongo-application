import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
    vus: 100, // Number of virtual users
    duration: "30s", // Duration of the test
};

export default function () {
    let res = http.get("http://127.0.0.1:1323/alive");

    check(res, {
        success: (res) => res.status == 200,
        succesString: (res) => res.body.trim() == "yes"
    });

    sleep(1);
}