<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; justify-content: space-between; align-items: center;">
          <span style="font-size: 18px; font-weight: bold;">
            {{ isEdit ? '编辑商品' : '新增商品' }}
          </span>
          <div>
            <el-button @click="handleCancel">取消</el-button>
            <el-button type="primary" @click="handleSave" :loading="submitting">
              保存
            </el-button>
            <el-button v-if="isEdit" type="success" @click="handleSaveAndContinue">
              保存并继续
            </el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="120px"
        v-loading="loading"
      >
        <el-row :gutter="20">
          <!-- 基本信息 -->
          <el-col :span="16">
            <el-card header="基本信息" style="margin-bottom: 20px;">
              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="商品名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入商品名称" />
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="商品编码" prop="code">
                    <el-input v-model="formData.code" placeholder="请输入商品编码" />
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row :gutter="20">
                <el-col :span="12">
                  <el-form-item label="商品分类" prop="categoryId">
                    <el-tree-select
                      v-model="formData.categoryId"
                      :data="categoryTree"
                      :props="{ label: 'name', value: 'id', children: 'children' }"
                      placeholder="请选择商品分类"
                      clearable
                      check-strictly
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="12">
                  <el-form-item label="商品状态" prop="status">
                    <el-radio-group v-model="formData.status">
                      <el-radio value="draft">草稿</el-radio>
                      <el-radio value="active">上架</el-radio>
                      <el-radio value="inactive">下架</el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-col>
              </el-row>

              <el-form-item label="商品描述" prop="description">
                <el-input
                  v-model="formData.description"
                  type="textarea"
                  :rows="4"
                  placeholder="请输入商品描述"
                />
              </el-form-item>
            </el-card>

            <!-- 价格与库存 -->
            <el-card header="价格与库存" style="margin-bottom: 20px;">
              <el-row :gutter="20">
                <el-col :span="8">
                  <el-form-item label="销售价格" prop="price">
                    <el-input-number
                      v-model="formData.price"
                      :min="0"
                      :precision="2"
                      style="width: 100%;"
                      placeholder="销售价格"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8">
                  <el-form-item label="成本价格" prop="costPrice">
                    <el-input-number
                      v-model="formData.costPrice"
                      :min="0"
                      :precision="2"
                      style="width: 100%;"
                      placeholder="成本价格"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8">
                  <el-form-item label="库存数量" prop="stock">
                    <el-input-number
                      v-model="formData.stock"
                      :min="0"
                      style="width: 100%;"
                      placeholder="库存数量"
                    />
                  </el-form-item>
                </el-col>
              </el-row>

              <el-row :gutter="20">
                <el-col :span="8">
                  <el-form-item label="预警数量" prop="alertQuantity">
                    <el-input-number
                      v-model="formData.alertQuantity"
                      :min="0"
                      style="width: 100%;"
                      placeholder="预警数量"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8">
                  <el-form-item label="重量(kg)" prop="weight">
                    <el-input-number
                      v-model="formData.weight"
                      :min="0"
                      :precision="2"
                      style="width: 100%;"
                      placeholder="重量"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="8">
                  <el-form-item label="体积(m³)" prop="volume">
                    <el-input-number
                      v-model="formData.volume"
                      :min="0"
                      :precision="3"
                      style="width: 100%;"
                      placeholder="体积"
                    />
                  </el-form-item>
                </el-col>
              </el-row>
            </el-card>

            <!-- SKU管理 -->
            <el-card header="规格管理">
              <div style="margin-bottom: 15px;">
                <el-button type="primary" size="small" @click="handleAddSKU">
                  <el-icon><Plus /></el-icon>
                  添加规格
                </el-button>
              </div>

              <el-table :data="formData.skus" border>
                <el-table-column prop="skuCode" label="SKU编码" width="150">
                  <template #default="{ row, $index }">
                    <el-input v-model="row.skuCode" placeholder="SKU编码" size="small" />
                  </template>
                </el-table-column>
                <el-table-column label="规格属性" min-width="200">
                  <template #default="{ row }">
                    <div v-for="(value, key) in row.attributes" :key="key" class="sku-attribute">
                      <span class="attr-key">{{ key }}:</span>
                      <el-input v-model="row.attributes[key]" size="small" style="width: 80px;" />
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="price" label="价格" width="120">
                  <template #default="{ row }">
                    <el-input-number
                      v-model="row.price"
                      :min="0"
                      :precision="2"
                      size="small"
                      style="width: 100%;"
                    />
                  </template>
                </el-table-column>
                <el-table-column prop="stock" label="库存" width="100">
                  <template #default="{ row }">
                    <el-input-number
                      v-model="row.stock"
                      :min="0"
                      size="small"
                      style="width: 100%;"
                    />
                  </template>
                </el-table-column>
                <el-table-column label="操作" width="100">
                  <template #default="{ $index }">
                    <el-button size="small" type="danger" @click="handleRemoveSKU($index)">
                      删除
                    </el-button>
                  </template>
                </el-table-column>
              </el-table>
            </el-card>
          </el-col>

          <!-- 图片上传 -->
          <el-col :span="8">
            <el-card header="商品图片" style="margin-bottom: 20px;">
              <el-upload
                v-model:file-list="imageList"
                :action="uploadUrl"
                :headers="uploadHeaders"
                :on-success="handleImageSuccess"
                :on-error="handleImageError"
                :before-upload="beforeImageUpload"
                list-type="picture-card"
                :limit="9"
                :on-exceed="handleExceed"
              >
                <el-icon><Plus /></el-icon>
              </el-upload>
              <div style="color: #999; font-size: 12px; margin-top: 10px;">
                支持jpg、png格式，单张图片不超过2MB，最多上传9张
              </div>
            </el-card>

            <!-- 商品属性 -->
            <el-card header="商品属性">
              <div style="margin-bottom: 15px;">
                <el-button size="small" @click="handleAddAttribute">
                  <el-icon><Plus /></el-icon>
                  添加属性
                </el-button>
              </div>

              <div v-for="(attr, index) in formData.attributes" :key="index" class="attribute-item">
                <el-row :gutter="10">
                  <el-col :span="8">
                    <el-input v-model="attr.name" placeholder="属性名" size="small" />
                  </el-col>
                  <el-col :span="12">
                    <el-input v-model="attr.value" placeholder="属性值" size="small" />
                  </el-col>
                  <el-col :span="4">
                    <el-button size="small" type="danger" @click="handleRemoveAttribute(index)">
                      删除
                    </el-button>
                  </el-col>
                </el-row>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, UploadProps, UploadUserFile } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { 
  getProduct, 
  createProduct, 
  updateProduct,
  getCategoryTree 
} from '@/api/product'

const route = useRoute()
const router = useRouter()

// 响应式数据
const loading = ref(false)
const submitting = ref(false)
const formRef = ref()
const categoryTree = ref([])
const imageList = ref<UploadUserFile[]>([])

// 表单数据
const formData = reactive({
  id: '',
  name: '',
  code: '',
  categoryId: '',
  status: 'draft',
  description: '',
  price: 0,
  costPrice: 0,
  stock: 0,
  alertQuantity: 10,
  weight: 0,
  volume: 0,
  images: [] as string[],
  skus: [] as any[],
  attributes: [] as any[]
})

// 表单验证规则
const formRules = {
  name: [
    { required: true, message: '请输入商品名称', trigger: 'blur' },
    { min: 1, max: 100, message: '商品名称长度在1到100个字符', trigger: 'blur' }
  ],
  code: [
    { required: true, message: '请输入商品编码', trigger: 'blur' },
    { min: 1, max: 50, message: '商品编码长度在1到50个字符', trigger: 'blur' }
  ],
  categoryId: [
    { required: true, message: '请选择商品分类', trigger: 'change' }
  ],
  price: [
    { required: true, message: '请输入销售价格', trigger: 'blur' },
    { type: 'number', min: 0, message: '销售价格不能小于0', trigger: 'blur' }
  ]
}

// 计算属性
const isEdit = computed(() => !!route.params.id)

// 上传配置
const uploadUrl = '/api/upload/images'
const uploadHeaders = {
  'Authorization': `Bearer ${localStorage.getItem('token')}`
}

// 方法
const loadCategories = async () => {
  try {
    const response = await getCategoryTree()
    categoryTree.value = response.data.data || []
  } catch (error) {
    console.error('加载分类失败:', error)
  }
}

const loadProduct = async () => {
  if (!isEdit.value) return
  
  loading.value = true
  try {
    const response = await getProduct(route.params.id as string)
    const product = response.data.data
    
    Object.assign(formData, {
      id: product.id,
      name: product.name,
      code: product.code,
      categoryId: product.category_id,
      status: product.status,
      description: product.description || '',
      price: product.price,
      costPrice: product.cost_price || 0,
      stock: product.stock || 0,
      alertQuantity: product.alert_quantity || 10,
      weight: product.weight || 0,
      volume: product.volume || 0,
      images: product.images || [],
      skus: product.skus || [],
      attributes: product.attributes || []
    })
    
    // 设置图片列表
    imageList.value = (product.images || []).map((url: string, index: number) => ({
      name: `image-${index}`,
      url: url,
      status: 'success'
    }))
  } catch (error) {
    ElMessage.error('加载商品信息失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const handleAddSKU = () => {
  formData.skus.push({
    skuCode: '',
    attributes: {},
    price: formData.price,
    stock: 0
  })
}

const handleRemoveSKU = (index: number) => {
  formData.skus.splice(index, 1)
}

const handleAddAttribute = () => {
  formData.attributes.push({
    name: '',
    value: ''
  })
}

const handleRemoveAttribute = (index: number) => {
  formData.attributes.splice(index, 1)
}

const handleImageSuccess: UploadProps['onSuccess'] = (response, file) => {
  if (response.code === 200) {
    formData.images.push(response.data.url)
    ElMessage.success('图片上传成功')
  } else {
    ElMessage.error(response.message || '图片上传失败')
  }
}

const handleImageError: UploadProps['onError'] = (error) => {
  ElMessage.error('图片上传失败')
  console.error(error)
}

const beforeImageUpload: UploadProps['beforeUpload'] = (file) => {
  const isImage = file.type.startsWith('image/')
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!')
    return false
  }
  return true
}

const handleExceed: UploadProps['onExceed'] = () => {
  ElMessage.warning('最多只能上传9张图片')
}

const handleSave = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    submitting.value = true
    
    const data = {
      ...formData,
      images: formData.images,
      skus: formData.skus.filter(sku => sku.skuCode),
      attributes: formData.attributes.filter(attr => attr.name && attr.value)
    }
    
    if (isEdit.value) {
      await updateProduct(formData.id, data)
      ElMessage.success('商品更新成功')
    } else {
      await createProduct(data)
      ElMessage.success('商品创建成功')
    }
    
    router.push('/products')
  } catch (error) {
    if (error !== false) {
      ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
      console.error(error)
    }
  } finally {
    submitting.value = false
  }
}

const handleSaveAndContinue = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    submitting.value = true
    
    const data = {
      ...formData,
      images: formData.images,
      skus: formData.skus.filter(sku => sku.skuCode),
      attributes: formData.attributes.filter(attr => attr.name && attr.value)
    }
    
    if (isEdit.value) {
      await updateProduct(formData.id, data)
      ElMessage.success('商品更新成功')
    } else {
      const response = await createProduct(data)
      ElMessage.success('商品创建成功')
      // 跳转到编辑页面
      router.push(`/products/edit/${response.data.data.id}`)
    }
  } catch (error) {
    if (error !== false) {
      ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
      console.error(error)
    }
  } finally {
    submitting.value = false
  }
}

const handleCancel = () => {
  router.push('/products')
}

// 生命周期
onMounted(() => {
  loadCategories()
  loadProduct()
})
</script>

<style scoped>
.sku-attribute {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.attr-key {
  min-width: 60px;
  font-size: 12px;
  color: #666;
}

.attribute-item {
  margin-bottom: 10px;
}
</style>
