#include <stdio.h>

int g_count = 0;

void Merge(int A[], int p, int q, int r)
{
    int n1 = q - p + 1;
    int n2 = r - q;
    int i, j, L[n1 + 1], R[n2 + 1];
    for (int i = 0; i < n1; ++i)
        L[i] = A[p + i];
    L[n1] = 1000000;
    for (int j = 0; j < n2; ++j)
        R[j] = A[q + j + 1];
    R[n2] = 1000000;
    i = j = 0;
    while (p <= r)
    {
        if (L[i] <= R[j])
            A[p++] = L[i++];
        else
        {
            A[p++] = R[j++];
            g_count += (q + 1) - p - i;
        }
    }
}

void MergeSort(int A[], int p, int r)
{
    int q;
    if (p < r)
    {
        q = (p + r) / 2;
        MergeSort(A, p, q);
        MergeSort(A, q + 1, r);
        Merge(A, p, q, r);
    }
}

int main()
{
    int A[9] = {5, 4, 1, 8, 3, 11, 7, 2, 9};
    g_count = 0;
    MergeSort(A, 0, sizeof(A));
    printf("%d", g_count);
    return 0;
}