-- Initial data only. The models in models/models.go create the tables via
-- database.DB.AutoMigrate in bootstrap.NewApp. Import this file into the
-- database configured in common/config/config.yml after that first startup.
-- The password for all three demo users is: Admin@2026!
INSERT INTO casbin_demo_users (username, password, role) VALUES
    ('admin', '$2a$10$Cs.5wXXiKntpN7H/SJGsLezrW1efwrr2WlFlz1dnpZ6VcJmKB2dB.', 'admin'),
    ('operate', '$2a$10$Cs.5wXXiKntpN7H/SJGsLezrW1efwrr2WlFlz1dnpZ6VcJmKB2dB.', 'operate'),
    ('user', '$2a$10$Cs.5wXXiKntpN7H/SJGsLezrW1efwrr2WlFlz1dnpZ6VcJmKB2dB.', 'user')
ON DUPLICATE KEY UPDATE password = VALUES(password), role = VALUES(role);

-- admin: CRUD; operate: create/read/update; user: read only.
INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
    ('p', 'admin',   '/api/tests',     '(GET|POST)','','',''),
    ('p', 'admin',   '/api/tests/:id', '(GET|PUT|DELETE)','','',''),
    ('p', 'operate', '/api/tests',     '(GET|POST)','','',''),
    ('p', 'operate', '/api/tests/:id', '(GET|PUT)','','',''),
    ('p', 'user',    '/api/tests',     'GET','','',''),
    ('p', 'user',    '/api/tests/:id', 'GET','','','');

-- 20 sample records for the protected test CRUD API.
INSERT INTO casbin_demo_tests (id, name, content, owner) VALUES
    (1,  '测试数据 01', 'Casbin + JWT + Fiber 示例数据 01', 'admin'),
    (2,  '测试数据 02', 'Casbin + JWT + Fiber 示例数据 02', 'operate'),
    (3,  '测试数据 03', 'Casbin + JWT + Fiber 示例数据 03', 'user'),
    (4,  '测试数据 04', 'Casbin + JWT + Fiber 示例数据 04', 'admin'),
    (5,  '测试数据 05', 'Casbin + JWT + Fiber 示例数据 05', 'operate'),
    (6,  '测试数据 06', 'Casbin + JWT + Fiber 示例数据 06', 'user'),
    (7,  '测试数据 07', 'Casbin + JWT + Fiber 示例数据 07', 'admin'),
    (8,  '测试数据 08', 'Casbin + JWT + Fiber 示例数据 08', 'operate'),
    (9,  '测试数据 09', 'Casbin + JWT + Fiber 示例数据 09', 'user'),
    (10, '测试数据 10', 'Casbin + JWT + Fiber 示例数据 10', 'admin'),
    (11, '测试数据 11', 'Casbin + JWT + Fiber 示例数据 11', 'operate'),
    (12, '测试数据 12', 'Casbin + JWT + Fiber 示例数据 12', 'user'),
    (13, '测试数据 13', 'Casbin + JWT + Fiber 示例数据 13', 'admin'),
    (14, '测试数据 14', 'Casbin + JWT + Fiber 示例数据 14', 'operate'),
    (15, '测试数据 15', 'Casbin + JWT + Fiber 示例数据 15', 'user'),
    (16, '测试数据 16', 'Casbin + JWT + Fiber 示例数据 16', 'admin'),
    (17, '测试数据 17', 'Casbin + JWT + Fiber 示例数据 17', 'operate'),
    (18, '测试数据 18', 'Casbin + JWT + Fiber 示例数据 18', 'user'),
    (19, '测试数据 19', 'Casbin + JWT + Fiber 示例数据 19', 'admin'),
    (20, '测试数据 20', 'Casbin + JWT + Fiber 示例数据 20', 'operate')
ON DUPLICATE KEY UPDATE
    name = VALUES(name), content = VALUES(content), owner = VALUES(owner);
