//Bismillah Hir Rahman Nir Raheem ( بسم الله الرحمن الرحيم )
#include <bits/stdc++.h>
using namespace std;

int main()
{
    long long n,a[1000],ans=0,max;
    scanf("%I64d",&n);
    for(int i=0;i<n;i++)
        scanf("%I64d",&a[i]);
    while(1)
    {
        max=0;
        for(int i=0;i<n;i++){
            if(a[max]<=a[i])
                max=i;
        }
        if(max==0)
            break;
        a[0]++;
        a[max]--;
        ans++;
    }
    printf("%I64d\n",ans);

    return 0;
}
