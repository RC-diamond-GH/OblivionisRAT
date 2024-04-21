import { invoke } from '@tauri-apps/api';

export function setupCounter(element: HTMLButtonElement) {
  let counter = 0, str = '';
  const setCounter = (count: number) => {
    invoke('hello', { name: 'world'}).then((res) => {
      console.log(res);
      str = res as string;
    })
    counter = count
    element.innerHTML = `count is ${counter}, str is ${str}`
  }
  element.addEventListener('click', () => setCounter(counter + 1))
  setCounter(0)
}
