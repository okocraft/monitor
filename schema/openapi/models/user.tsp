import "./uuid.tsp";
import "./role.tsp";

model User {
  id: UUID;
  nickname: string;
  last_access: utcDateTime;
  created_at: utcDateTime;
  updated_at: utcDateTime;
  role: Role;
}
