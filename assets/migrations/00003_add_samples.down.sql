DELETE FROM roles
WHERE
    name IN ('Admin', 'Default');

DELETE FROM model_types
WHERE
    type IN (
        'regionsegmentation',
        'linesegmentation',
        'textrecognition'
    );

DELETE FROM models
WHERE
    name IN (
        'yolov9-lines-within-regions-1',
        'TrOCR-norhand-v3'
    );
