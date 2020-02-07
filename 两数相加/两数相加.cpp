//#include<stdio.h> 
//#include<stdlib.h> 
////Definition for singly-linked list.
//struct ListNode {
//	int val;
//	struct ListNode *next;
//};
//
//struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
//	printf("%d", l1->val);
//	struct ListNode* temp = l1;
//	for(int i=0; temp->next==NULL; i++){
//		printf("%d\n", temp->val);
//		temp = temp->next; 
//	}
//}
//
//int main(){
//	printf("hello\n");
//	struct ListNode* a, *b, res;
//	a->val = 2;
//	a->next->val = 4;
//	a->next->next->val = 3;
//	b->val = 5;
//	b->next->val = 6;
//	b->next->next->val = 4;
//	printf("%d", a->val);
////	printf("%d", a.val);
//	addTwoNumbers(a, b);
//}


#include<stdio.h>
#include<stdlib.h>

struct ListNode {
	int val;
	struct ListNode *next;
};

int output(struct ListNode *L){
	struct ListNode *p;
	for(p = L->next; p!=NULL; p = p->next){
		printf("%d",p->val);
	}
	printf("\n");
}

struct ListNode *Create(int arr[], struct ListNode *L){
	int i;
	struct ListNode *r;
	r = L;
	for(i=0; i<3; i++){
		struct ListNode *p;
		p = (struct ListNode *)malloc(sizeof(struct ListNode));
		p->val = arr[i];
		L->next = p;
		L = p;
		printf("arr[%d] = %d\n", i, arr[i]);
	}
	L->next = NULL;
	return L;
}


int main(){
	struct ListNode *a, *L;
	L = (struct ListNode *)malloc(sizeof(struct ListNode));
	if (L == NULL)
		printf("error");
	L->next = NULL;
	int i;
//	for(i=0; i < 9; i++){
//		struct ListNode *p;
//		p = (struct ListNode *)malloc(sizeof(struct ListNode));
//		a->next = p;
//		a->val = i;
//		a =
//		printf("%d\n", i);
//	}
	a = L;
	i = 10;
	while(i--){
		struct ListNode *p;
		p = (struct ListNode *)malloc(sizeof(struct ListNode));
		p->val = i;
		a->next = p;
		a = p;
		printf("%d\n", i);
	}
	a->next = NULL;
//	printf("hello");
//	a->next->val = 2;
//	a->next->next->val = 3;
	struct ListNode *p;
	for (p = L->next; p != NULL; p = p->next)
	{
		printf("%d", p->val);
	}

	
	int a_1[] = {1,2,3};
	int a_2[] = {2,3,4};
	struct ListNode *n, *m;
	n = (struct ListNode *)malloc(sizeof(struct ListNode));
	if (n == NULL)
		printf("error");
	n->next = NULL;
	
	m = (struct ListNode *)malloc(sizeof(struct ListNode));
	if (m == NULL)
		printf("error");
	m->next = NULL;
	
	n = Create(a_1, n);
	m = Create(a_2, m);
	
	output(n);
	output(m);
	
	return 0; 
}

