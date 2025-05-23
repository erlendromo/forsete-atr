INSERT INTO
    "role" (name)
VALUES
    ('Admin'),
    ('Default');

INSERT INTO
    "model_type" (type)
VALUES
    ('regionsegmentation'),
    ('linesegmentation'),
    ('textrecognition');

INSERT INTO
    "model" (name, path, model_type_id)
VALUES
    (
        'yolov9-lines-within-regions-1',
        'assets/models/linesegmentation/yolov9-lines-within-regions-1/model.pt',
        2
    ),
    (
        'TrOCR-norhand-v3',
        'assets/models/textrecognition/TrOCR-norhand-v3',
        3
    );
