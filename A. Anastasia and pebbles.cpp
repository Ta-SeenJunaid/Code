//Bismillah Hir Rahman Nir Raheem ( بسم الله الرحمن الرحيم )
#include <bits/stdc++.h>
using namespace std;



int main()
{
    long long n,k,a[1000],c,sum=0;
    scanf("%I64d %I64d",&n,&k);

    for(int i=0;i<n;i++){
        scanf("%I64d",&a[i]);
        sum=sum+a[i];
    }
   // cout<<sum;
    printf("%I64d\n",sum/n);

   /* for(int i=0;i<n;i++){
            if(a[i]<n)

    }*/


    return 0;
}
