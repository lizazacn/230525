<template>
  <div style="width: 300px">
    <el-form ref="formRef" :model="nameValidateForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="姓名" prop="name" :rules="[
        { required: true, message: 'name is required' },
      ]">
        <el-input v-model.number="nameValidateForm.name" type="text" autocomplete="off" @input="handleInputChange" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(formRef)">Submit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, inject } from 'vue'
import type { FormInstance } from 'element-plus'
import { useRouter } from 'vue-router'


const socket = inject('socket') as any;

const id = inject('id')


const formRef = ref<FormInstance>()
const router = useRouter()

const nameValidateForm = reactive({
  name: '',
})

const handleInputChange = (value: string | number) => {


  console.log((window as any).id, 377777);

  const obj = {
    name: value,
    id: (window as any).id
  }

  socket.emit('update', JSON.stringify(obj))
}

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid: any) => {
    if (valid) {
      console.log('submit!')
      router.push({ name: 'page2' })
    } else {
      console.log('error submit!')
      return false
    }
  })
}

</script>
