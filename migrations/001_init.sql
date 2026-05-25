CREATE TABLE IF NOT EXISTS apps (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS screens (
    id VARCHAR(50) PRIMARY KEY,
    app_id VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    CONSTRAINT fk_screens_app FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS elements (
    id VARCHAR(50) PRIMARY KEY,
    screen_id VARCHAR(50) NOT NULL,
    parent_id VARCHAR(50),
    name VARCHAR(100) NOT NULL,
    type VARCHAR(30) NOT NULL,
    selector VARCHAR(255) NOT NULL,
    text_label VARCHAR(100),
    position JSON,
    visible_condition JSON,
    CONSTRAINT fk_elements_screen FOREIGN KEY (screen_id) REFERENCES screens(id) ON DELETE CASCADE,
    CONSTRAINT fk_elements_parent FOREIGN KEY (parent_id) REFERENCES elements(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS actions (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(30) NOT NULL,
    target_element_id VARCHAR(50),
    parameters JSON,
    preconditions JSON,
    expected_outcome JSON,
    priority INT DEFAULT 1,
    CONSTRAINT fk_actions_element FOREIGN KEY (target_element_id) REFERENCES elements(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS goals (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS goal_action_mapping (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    goal_id VARCHAR(50) NOT NULL,
    action_id VARCHAR(50) NOT NULL,
    step_order INT NOT NULL,
    context_condition JSON,
    CONSTRAINT fk_mapping_goal FOREIGN KEY (goal_id) REFERENCES goals(id) ON DELETE CASCADE,
    CONSTRAINT fk_mapping_action FOREIGN KEY (action_id) REFERENCES actions(id) ON DELETE CASCADE,
    CONSTRAINT uq_goal_step UNIQUE (goal_id, step_order)
);
