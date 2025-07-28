#include <memory>
#include <limits>
#include <iostream>

using std::shared_ptr;

struct Node {
    
    Node(int value) {
        this->value = value;
        this->next =nullptr; 
        this->prev = nullptr;
    }
    
    Node() {
        this->value = std::numeric_limits<int>::max(); 
        this->next = nullptr;
        this->prev = nullptr;
    }

    int value;
    shared_ptr<Node> next;
    shared_ptr<Node> prev;
};

class MyLinkedList {
    private:
    shared_ptr<Node> head;
    shared_ptr<Node> tail;
    int len;
    
    public:
    MyLinkedList()  {
        // head is invalid node
        this->head = std::make_shared<Node>();
        this->tail = this->head;
        this->len = 0;
    }
    
    int get(int index) {
        if (index >= this->len || index < 0) {
            return -1;
        }
        
        auto cursor = this->head;
        for (auto i = 0;i <= index; ++i) {
            cursor = cursor->next;
        }
        return cursor->value;
    }
    
    void addAtHead(int val) {
        auto tmp = std::make_shared<Node>(val);

        // head-->b
        // --> head --> a --> b
        tmp->next = this->head->next;
        // need to judge this node that is the first inserted node;
        if (this->head->next != nullptr) {
            this->head->next->prev = tmp;
        } else {
            this->tail = tmp;
        }

        tmp->prev = head;
        head->next = tmp;
        this->len += 1;
    }
    
    void addAtTail(int val) {
        auto tmp = std::make_shared<Node>(val);
        this->tail->next = tmp;
        tmp->prev = this->tail;
        this->tail = tmp;
        this->len += 1;
    }
    
    void addAtIndex(int index, int val) {   
        if (index == this->len) {
            this->addAtTail(val);
            return;
        }

        if (index > this->len || index < 0) {
            return;
        }

        auto tmpNode = std::make_shared<Node>(val);
        auto cursor = this->head;
        for (auto i = 0; i <= index; ++i) {
            cursor = cursor->next;
        }

        auto prev = cursor->prev;
        tmpNode->prev = prev;
        prev->next = tmpNode;
        tmpNode->next = cursor;
        cursor->prev = tmpNode;

        this->len += 1;
    }
    
    void deleteAtIndex(int index) {
        if (index >= this->len || index < 0) {
            return;
        }

        if (index == this->len - 1) {
            auto prev = this->tail->prev;
            this->tail->next = this->tail->prev = nullptr;
            this->tail = prev;
            this->tail->next = nullptr;
            this->len -= 1;
            return;
        }

        auto cursor = this->head;
        for(auto i = 0; i <= index; ++i) {
            cursor = cursor->next;
        }

        std::cout << "delete:" << cursor->value << std::endl;

        // a->b->c ==> a->c
        auto prev = cursor->prev;
        auto next = cursor->next;
        cursor->next = cursor->prev = nullptr;
        prev->next = next;
        next->prev = prev;
        this->len -= 1;

        return;
    }
};