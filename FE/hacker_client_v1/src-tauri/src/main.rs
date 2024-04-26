// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

// 导入命令模块
mod command;

fn main() {
  tauri::Builder::default()
    .invoke_handler(tauri::generate_handler![command::hello])
    .run(tauri::generate_context!())
    .expect("error while running tauri application");
}
