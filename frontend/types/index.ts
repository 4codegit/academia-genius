export interface User {
  id: number; username: string; email: string; full_name: string;
  role: 'student' | 'teacher' | 'admin'; created_at: string;
}
export interface LoginRequest { email: string; password: string }
export interface RegisterRequest { username: string; email: string; password: string; full_name: string }
export interface AuthResponse { token: string; user: User }
export interface News {
  id: number; title: string; slug: string; summary: string; content: string;
  image_url: string; published_at: string; is_active: boolean;
}
export type PhysicsTopic = 'Механика'|'МКТ'|'Термодинамика'|'Электростатика'|'Магнетизм'|'Оптика'|'СТО'|'Квантовая'
export type Difficulty = 'easy'|'medium'|'hard'|'olympiad'
export interface Problem {
  id: number; title: string; topic: PhysicsTopic; difficulty: Difficulty;
  content: string; solution: string; image_url: string; created_at: string;
}
export interface Course {
  id: number; title: string; description: string; instructor: string;
  image_url: string; price: number; duration: string; is_active: boolean;
}
export type BookCategory = 'Учебники'|'Задачники'|'Справочники'|'Монографии'|'Подготовка к олимпиадам'
export interface Book {
  id: number; title: string; author: string; category: BookCategory;
  description: string; cover_url: string; year: number; download_url: string;
}
export interface Alumnus {
  id: number; full_name: string; bio: string; photo_url: string;
  graduation_year: number; university: string; is_featured: boolean; sort_order: number;
}
export interface AlumniAward {
  id: number; alumni_id: number; award_title: string; competition: string;
  year: number; description: string; sort_order: number;
}
export interface FeaturedAlumnus extends Alumnus { awards: AlumniAward[] }
export interface Webinar {
  id: number; title: string; description: string; speaker: string;
  event_date: string; duration_min: number; platform_url: string; is_active: boolean;
}
export interface KnowledgeMapEntry { topic: PhysicsTopic; progress: number }
export interface UserStats {
  total_solved: number; by_topic: Record<string, number>;
  by_difficulty: Record<string, number>; knowledge_map: KnowledgeMapEntry[];
  streak_days: number; last_active: string;
}
export interface PaginatedResponse<T> { data: T[]; total: number; page: number; limit: number; total_pages: number }
