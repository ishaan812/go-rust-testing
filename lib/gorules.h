typedef enum {
    Null,
    Bool,
    Number,
    String,
    Array,
    Object,
} ValueType;

typedef struct {
    ValueType type;
    union {
        int bool_value;
        double number_value;
        char* string_value;
        struct {
            struct Value* array_elements;
            size_t array_len;
        };
        struct {
            struct KeyValuePair* object_elements;
            size_t object_len;
        };
    } data;
} Value;

typedef struct {
    char* key;
    Value value;
} KeyValuePair;


void hello(char *name);
char* evaluate(char *decisionContent, char *params);
