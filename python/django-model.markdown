# Django Model

## aggregate(聚合)


代码片段：两字段相乘

```python
from django.db.models import F, FloatField
...
Order.objects.filter(<xxx>).aggregate(
    total=Sum(F('unit_price') * F('num')),
    output_field=FloatField()
)['total']
```

## annotation(注解)

对数据分组后，在聚合


## select_related

## prefetch_related