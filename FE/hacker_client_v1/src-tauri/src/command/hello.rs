#[tauri::command]
pub fn hello(name: &str) -> String {
  format!("Hello, {} from Rust!", name)
}