DELETE FROM "model"
WHERE
    name IN (
        'yolov9-lines-within-regions-1',
        'TrOCR-norhand-v3'
    );

DELETE FROM "model_type"
WHERE
    type IN (
        'regionsegmentation',
        'linesegmentation',
        'textrecognition'
    );

DELETE FROM "role"
WHERE
    name IN ('Admin', 'Default');
