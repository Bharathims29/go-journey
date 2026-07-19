CREATE TABLE IF NOT EXISTS student_marks (
	id BIGSERIAL PRIMARY KEY,
	student_name TEXT NOT NULL,
	roll_number TEXT NOT NULL,
	subject TEXT NOT NULL,
	marks INTEGER NOT NULL CHECK (marks >= 0),
	max_marks INTEGER NOT NULL CHECK (max_marks > 0),
	grade TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_student_marks_roll_number
ON student_marks (roll_number);
