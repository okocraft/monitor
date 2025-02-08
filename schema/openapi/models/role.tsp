import "./uuid.tsp";

model Role {
  id: UUID;
  name: string;
  priority: int32;
  created_at: utcDateTime;
  updated_at: utcDateTime;
}
