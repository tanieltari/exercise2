import {
  debounceTime,
  distinctUntilChanged,
  filter,
  fromEvent,
  map,
  switchMap,
} from "rxjs";
import { fromFetch } from "rxjs/fetch";

const query = document.getElementById("q") as HTMLInputElement | null;

const refresh = document.getElementById("refresh") as HTMLButtonElement | null;
const bus = document.getElementById("bus") as HTMLInputElement | null;
const train = document.getElementById("train") as HTMLInputElement | null;

fromEvent<InputEvent>(query!, "input")
  .pipe(
    map((e: any) => e.target.value),
    map((value: string) => value.trim().toLowerCase()),
    debounceTime(500),
    filter((term) => term.length > 0),
    distinctUntilChanged(),
    map((term) => `http://localhost:4000/stops?q=${term}`),
    switchMap((url) => fromFetch(url).pipe(switchMap((r) => r.json())))
  )
  .subscribe((x) => console.log(x));
